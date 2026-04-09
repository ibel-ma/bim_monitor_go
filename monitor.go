package main

import (
	"fmt"
	"time"
)

func LiveMonitor(name, lid string, ref int, filter string) {
	departures, _ := GetDepartures(lid)
	// TODO filter departures for direction
	last := time.Now()

	// filter direction
	departures, _ = filterDepartures(departures, filter)

	ClearTerminal()
	PrintDeparturesLive(name, departures, last)

	for {
		time.Sleep(time.Minute)
		// Decrease departure countdown
		for i := range departures {
			departures[i].Countdown--
		}
		// Remove past departures
		fil := departures[:0]
		for _, d := range departures {
			if d.Countdown >= 0 {
				fil = append(fil, d)
			}
		}
		departures = fil

		// Refresh departures if the ref time has passed
		if time.Since(last) >= time.Duration(ref)*time.Minute {
			newDepartures, _ := GetDepartures(lid)
			departures, _ = filterDepartures(newDepartures, filter)
			last = time.Now()
		}

		ClearTerminal()
		PrintDeparturesLive(name, departures, last)
	}
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
