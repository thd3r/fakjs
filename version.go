package main

import (
	"encoding/json"
	"io"
)

var CurrentVersion = "v1.1.2"

func Version() string {
	client := NewClient()

	resp, err := client.Do("GET", "https://api.github.com/repos/thd3r/fakjs/releases/latest")
	if err != nil {
		return CurrentVersion + " " + ColoredText("magenta", "unknown")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CurrentVersion + " " + ColoredText("magenta", "unknown")
	}

	var dataRelease = struct {
		ReleaseVersion string `json:"tag_name"`
	}{}

	if err := json.Unmarshal(body, &dataRelease); err != nil {
		return CurrentVersion + " " + ColoredText("magenta", "unknown")
	}

	if CurrentVersion < dataRelease.ReleaseVersion {
		return CurrentVersion + " " + ColoredText("red", "outdated")
	}
	if CurrentVersion == dataRelease.ReleaseVersion {
		return CurrentVersion + " " + ColoredText("green", "latest")
	}

	return CurrentVersion + " " + ColoredText("magenta", "unknown")
}
