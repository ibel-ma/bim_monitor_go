package main

import (
	"fmt"
	"strings"
	"time"
)

func PrintDepartures(name string, departures []Abfahrt) {
	now := time.Now()
	fmt.Println(strings.Repeat("=", 75))
	fmt.Println(" ", name)
	fmt.Println("  As of:", now.Format("15:04:05"))
	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf(" %-19s %-38s %-7s  %s\n", "Line", "Direction", "Time", "Status")
	fmt.Println(strings.Repeat("-", 75))
	// print departures
	for _, x := range departures {
		fmt.Printf(" %-12s %-6s %-38s (%s)  [%2d]\n", x.Type, x.Line, x.Direction, x.RealTime, x.Countdown)
	}
	fmt.Println(strings.Repeat("=", 75))
}

func print_help() {
	help_text :=
		`
        bim_monitor [ARGUMENTS]
        For one arguments the script outputs a departure board for the given location.
        For two arguments the script monitors the departures live. 
        Updates every minute and fetches new data from the API every 5 minutes.

        Arguments:
        arg1: Location name e.g. "Steyrergasse", "Jakominiplatz"
        arg2: Direction (Optional) depends on line, e.g. "Liebenau"

        Example:
		go build .
		./bim_monitor Location direction(optional) flag(optional)
        `
	fmt.Printf(help_text)
}

func PrintDeparturesLive(name string, departures []Abfahrt, last time.Time) {
	now := time.Now()

	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf(" %s\n", name)

	// print last update
	last_min := int(now.Sub(last).Minutes())
	fmt.Printf(" As of: %-46s Last update: %d min\n", now.Format("15:04"), last_min)
	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf(" %-19s %-38s %-7s  %s\n", "Line", "Direction", "Time", "Status")
	fmt.Println(strings.Repeat("-", 75))
	// print departures
	for _, x := range departures {
		//fmt.Printf(" %-6s %-38s (%s)  [%2d]\n", x.Line, x.Direction, x.RealTime, x.Countdown)
		fmt.Printf(" %-12s %-6s %-38s (%s)  [%2d]\n", x.Type, x.Line, x.Direction, x.RealTime, x.Countdown)
	}
	fmt.Println(strings.Repeat("=", 75))
	fmt.Println("Running... Press Ctrl+C to stop.")
}
