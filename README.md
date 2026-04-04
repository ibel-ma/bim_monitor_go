# bim_monitor

`bim_monitor` is a small go script for fetching real-time departure information for Graz public transport from the HAFAS API at `verkehrsauskunft.verbundlinie.at`.

## Overview

The script provides functions to:

- Search for stops by name
- Retrieve departure board for a stop
- Filter departure data for direction
- Live monitor departure board for a stop

## Functions

### `hafasRequest(req any) (map[string]any, error)`
Sends an HTTP POST request to the HAFAS API and returns the raw JSON response. It uses fixed client and authentication payload data.

### `SearchStops(name string) ([]StopSearchResult, error)`
Searches for stops matching the provided name and returns all matches.

Returns:
- A list of objects with `Name`, `Lid`, and `ExtID`

```go
type StopSearchResult struct {
	Name  string
	Lid   string
	ExtID string
}
```

### `Departures`
A dataclass representing a single departure entry.

```go
type Abfahrt struct {
	Line, Type, Direction, PlanTime, RealTime string
	Countdown                                 int
}
```

Fields:
- `Line`: Line number or name
- `Type`: Vehicle type (e.g. tram, city bus)
- `Direction`: Direction of travel
- `PlanTime`: Scheduled departure time
- `RealTime`: Real-time departure time (if available)
- `Countdown`: Minutes until departure

### `parseTime(timeStr string, dateStr string) (time.Time, error)`
Parses a HAFAS time string (`HHMM` and `YYYYMMDD`) into a `time.Time` object in the Vienna timezone.

### `GetDepartures(lid string) ([]Abfahrt, error)`
Loads the next departures for a stop and returns them as a list of `Abfahrt` objects.

### `PrintDepartures(name string, departures []Abfahrt)`
Prints a formatted departure board to the console.

- Uses `GetDepartures()` to fetch data
- Displays line, direction, time, delay status, and countdown

### `print_help()`
Prints help text

### `LiveMonitor(name, lid string, ref int, filter string)`
Prints departure board and updates every minute. Exit with `ctrl+C`

Arguments:
- `name`: Name of Stop
- `lid`: Location ID for Stop
- `ref`: minutes to update HAFAS request
- `filter`: Direction to filter departures

## Requirements

- go 1.22

## Usage

```bash
go build .
./bim_monitor
./graz-bim-monitor "Steyrergasse"
./graz-bim-monitor "Steyrergasse" --live
```

```
===========================================================================
  Steyrergasse
  As of: 18:51:49
===========================================================================
 Line                Direction                              Time     Status
---------------------------------------------------------------------------
 Straßenbahn  4      Liebenau                               (18:58)  [ 6]
 Straßenbahn  4      Liebenau                               (19:13)  [21]
 Straßenbahn  4      Liebenau                               (19:28)  [36]
===========================================================================
```

## Notes

This script uses the HAFAS API at `verkehrsauskunft.verbundlinie.at` and is designed for Graz Linien timetable data. Changes to the API or authentication may require updates.
The executable was build on macos(arm) and may not work on other systems.

## Thanks

Thanks Melissa for the great idea.
Also thanks to the team of [straba.at](https://straba.at/) for the inspiration.