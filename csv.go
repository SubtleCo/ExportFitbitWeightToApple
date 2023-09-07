package main

import (
	"fmt"
	"strings"
)

func ConvertToCSV(entries map[string]float64) string {
	header := "Date,Weight"
	var rows []string
	for date, weight := range entries {
		rows = append(rows, fmt.Sprintf("%s,%.2f", date, weight))
	}

	return header + "\n" + strings.Join(rows, "\n")
}
