package main

import (
	"fmt"
	"os"
)

func main() {

	// help
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		print_help()
		return
	}

	// arguments parsing
	name := os.Args[1]
	var direction, liveFlag string
	if len(os.Args) >= 3 {
		direction = os.Args[2]
	}
	if len(os.Args) >= 4 {
		liveFlag = os.Args[3]
	}

	// search for stops
	stops, err := SearchStops(name)
	if err != nil {
		fmt.Printf("Error searching for stops: %v\n", err)
		return
	}
	if len(stops) == 0 {
		fmt.Println("No stops found")
		return
	}
	LID := stops[0].Lid

	// request departures
	departures, err := GetDepartures(LID)
	if err != nil {
		fmt.Printf("Error requesting departures: %v\n", err)
		return
	}

	// filter directions
	if direction != "" {
		departures, err = filterDepartures(departures, direction)
		if err != nil {
			fmt.Printf("Error filtering departures: %v\n", err)
			return
		}
	}

	// print departure board or live monitoring
	if liveFlag == "--live" {
		min := 5
		LiveMonitor(name, LID, min, direction)
	} else {
		PrintDepartures(name, departures)
	}
}
