package main

import "fmt"

func ConvertToDigitalFormat(hour, minute, second int) string {
	digitalTime := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	return digitalTime
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	return hours, minutes, seconds
}
