package main

import (
	"encoding/json"
	"os"
)

type WeightEntry struct {
	LogID  int64   `json:"logId"`
	Weight float64 `json:"weight"`
	BMI    float64 `json:"bmi"`
	Fat    float64 `json:"Fat"`
	Date   string  `json:"date"`
	Time   string  `json:"time"`
	Source string  `json:"source"`
}

func LoadEntries(filename string) ([]WeightEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []WeightEntry
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&entries); err != nil {
		return nil, err
	}

	return entries, nil
}
