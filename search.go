package main

import (
	"errors"
)

type StopSearchResult struct {
	Name  string
	Lid   string
	ExtID string
}

func SearchStops(name string) ([]StopSearchResult, error) {
	//logger := slog.Default()
	if name == "" {
		return nil, errors.New("empty")
	}
	req := map[string]any{
		"meth": "LocMatch",
		"id":   "1|1|",
		"req": map[string]any{
			"input": map[string]any{
				"loc":    map[string]any{"type": "S", "name": name},
				"maxLoc": 10, "field": "S"}}}
	data, err := hafasRequest(req)
	if err != nil {
		return nil, err
	}
	svc := data["svcResL"].([]any)[0].(map[string]any)
	match := svc["res"].(map[string]any)["match"].(map[string]any)["locL"].([]any)
	res := []StopSearchResult{}
	for _, m := range match {
		mm := m.(map[string]any)
		res = append(res, StopSearchResult{mm["name"].(string), mm["lid"].(string), mm["extId"].(string)})
	}
	return res, nil
}
