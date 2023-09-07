package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func listWeightFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var weightFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "weight") && strings.HasSuffix(file.Name(), ".json") {
			weightFiles = append(weightFiles, path.Join(dir, file.Name()))
		}
	}
	return weightFiles, nil
}

func main() {
	// Check if directory is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program-name> <directory-path>")
		return
	}

	dir := os.Args[1] // Get the directory from CLI arguments

	files, err := listWeightFiles(dir)
	if err != nil {
		fmt.Printf("Failed to list files: %s\n", err)
	}

	allEntries := make(map[string]float64)

	for _, file := range files {
		entries, err := LoadEntries(file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, entry := range entries {
			// If multiple entries exist for the same date, average them
			if weight, ok := allEntries[entry.Date]; ok {
				allEntries[entry.Date] = (weight + entry.Weight) / 2.0
			} else {
				allEntries[entry.Date] = entry.Weight
			}
		}
	}

	csv := ConvertToCSV(allEntries)
	file, err := os.Create("export.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(csv)
	file.Sync()
}
