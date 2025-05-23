package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DataObj struct {
	Name   string   `json:"name"`
	Regex  string   `json:"regex"`
	Result []string `json:"results"`
}

type DataOutput struct {
	Info      string               `json:"info"`
	Version   string               `json:"version"`
	Timestamp time.Time            `json:"timestamp"`
	Data      []map[string]DataObj `json:"data"`
}

var FilePath = fmt.Sprintf("%s/fakjs-%v.json", os.TempDir(), time.Now().UnixNano())

func JsonReport(data chan FinalResults) error {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("%s: %v\n", ColoredText("red", "error"), err)
	}
	defer file.Close()

	var results DataOutput

	results.Info = "fakjs-output"
	results.Version = VERSION
	results.Timestamp = time.Now()

	for d := range data {
		results.Data = append(results.Data, map[string]DataObj{
			d.Url: {
				Name:   d.Name,
				Regex:  d.Regex,
				Result: d.DataOut,
			},
		})
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(results); err != nil {
		fmt.Printf("%s: %v\n", ColoredText("red", "error"), err)
	}

	return nil
}
