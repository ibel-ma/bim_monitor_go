package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Abfahrt struct {
	Line, Type, Direction, PlanTime, RealTime string
	Countdown                                 int
}

func parseTime(timeStr string, dateStr string) (time.Time, error) {
	if len(timeStr) < 4 {
		return time.Time{}, errors.New("bad")
	}
	// Example date string
	//date := "20261130" // YYYYMMDD

	// Parse the date string
	year, _ := strconv.Atoi(dateStr[0:4])
	monthIdx, _ := strconv.Atoi(dateStr[4:6])
	month := time.Month(monthIdx)
	day, _ := strconv.Atoi(dateStr[6:8])
	hour, _ := strconv.Atoi(timeStr[0:2])
	minute, _ := strconv.Atoi(timeStr[2:4])

	// return time.ParseInLocation("15:04", timeStr[0:2]+":"+timeStr[2:4], time.Local)

	// Create a time object for the departure time
	dp := time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	return dp, nil
}

func GetDepartures(lid string) ([]Abfahrt, error) {
	req := map[string]any{
		"meth": "StationBoard", "id": "1|1|",
		"req": map[string]any{
			"stbLoc": map[string]any{"lid": lid},
			"type":   "DEP", "sort": "PT", "maxJny": 10,
		},
	}
	data, err := hafasRequest(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %v", err)
	}

	// Safely extract response
	svcResL, ok := data["svcResL"].([]any)
	if !ok || len(svcResL) == 0 {
		return nil, errors.New("invalid API response format")
	}
	res, ok := svcResL[0].(map[string]any)["res"].(map[string]any)
	if !ok {
		return nil, errors.New("invalid API response format")
	}
	journeys, ok := res["jnyL"].([]any)
	if !ok {
		return nil, errors.New("invalid API response format")
	}

	product_list, ok := res["common"].(map[string]any)["prodL"].([]any)

	now := time.Now()
	out := []Abfahrt{}
	for _, j := range journeys {
		jj, ok := j.(map[string]any)
		if !ok {
			continue
		}
		stbStop, ok := jj["stbStop"].(map[string]any)
		if !ok {
			continue
		}
		date := jj["date"].(string) // e.g. "20261130"
		dTimeS, ok := stbStop["dTimeS"].(string)
		if !ok {
			continue
		}
		pd, err := parseTime(dTimeS, date)
		if err != nil {
			continue
		}

		var rd time.Time
		if v, ok := stbStop["dTimeR"]; ok {
			rs, ok := v.(string)
			if !ok {
				continue
			}
			rd, err = parseTime(rs, date)
			if err != nil {
				continue
			}
		}

		dep := pd
		if !rd.IsZero() {
			dep = rd
		}

		if dep.Before(now) {
			continue // Skip past departures
		}

		// Format real time if available, otherwise use planned time
		realTime := ""
		if !rd.IsZero() {
			realTime = rd.Format("15:04")
		} else {
			realTime = pd.Format("15:04")
		}

		// Get line and type from product list
		product_idx, ok := jj["prodX"].(float64)

		line := ""
		typ := ""
		if int(product_idx) < len(product_list) {
			prod := product_list[int(product_idx)].(map[string]any)
			line = prod["nameS"].(string)
			ctx := prod["prodCtx"].(map[string]any)
			typ = ctx["catOutL"].(string)
		}
		direction, ok := jj["dirTxt"].(string)
		out = append(out, Abfahrt{
			Line:      line,
			Type:      typ,
			Direction: direction,
			PlanTime:  pd.Format("15:04"),
			RealTime:  realTime,
			Countdown: int(dep.Sub(now).Minutes()),
		})
	}
	return out, nil
}

func filterDepartures(departures []Abfahrt, filter string) ([]Abfahrt, error) {
	// filter direction
	fil := departures[:0]
	for _, d := range departures {
		if strings.EqualFold(d.Direction, filter) {
			fil = append(fil, d)
		}
	}
	return fil, nil
}
