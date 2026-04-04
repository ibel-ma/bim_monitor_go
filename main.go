package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 || os.Args[1] == "--help" {
		print_help()
		return
	}

	// Argument-Parsing
	name := os.Args[1]
	var direction, liveFlag string
	if len(os.Args) >= 3 {
		direction = os.Args[2]
	}
	if len(os.Args) >= 4 {
		liveFlag = os.Args[3]
	}

	// Stop suchen
	stops, err := SearchStops(name)
	if err != nil {
		fmt.Printf("Fehler bei der Suche: %v\n", err)
		return
	}
	if len(stops) == 0 {
		fmt.Println("Keine Stopps gefunden")
		return
	}
	LID := stops[0].Lid

	// Abfahrten abrufen
	departures, err := GetDepartures(LID)
	if err != nil {
		fmt.Printf("Fehler beim Abrufen der Abfahrten: %v\n", err)
		return
	}

	// Logik für Richtung und Live-Modus
	if direction != "" {
		departures, err = filterDepartures(departures, direction)
		if err != nil {
			fmt.Printf("Fehler beim Filtern der Abfahrten: %v\n", err)
			return
		}
	}

	if liveFlag == "--live" {
		min := 5
		LiveMonitor(name, LID, min, direction)
	} else {
		PrintDepartures(name, departures)
	}
}
