package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DataObj struct {
	Target  string   `json:"target"`
	Name    string   `json:"name"`
	Regex   string   `json:"regex"`
	Results []string `json:"results"`
}

type DataOutput struct {
	Info      string    `json:"info"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Output    []DataObj `json:"data_output"`
}

var FilePath = fmt.Sprintf("%s/fakjs-%v.json", os.TempDir(), time.Now().UnixNano())

func JsonReport(verbose bool, data chan FinalResults) {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		FilteredVerboseOutput(verbose, fmt.Sprintf("%s: %v", ColoredText("red", "error"), err))
	}
	defer file.Close()

	var results DataOutput

	results.Info = "FakJs-Output"
	results.Version = CurrentVersion
	results.Timestamp = time.Now()

	for d := range data {
		results.Output = append(results.Output, DataObj{
			Target:  d.Target,
			Name:    d.Name,
			Regex:   d.Regex,
			Results: d.DataOut,
		})
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(results); err != nil {
		FilteredVerboseOutput(verbose, fmt.Sprintf("%s: %v", ColoredText("red", "error"), err))
	}

	if len(results.Output) > 0 {
		fmt.Printf(":: Report saved to %s\n", FilePath)
	}
}
