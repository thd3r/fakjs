package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed configs.yaml
var embeddedConfigs []byte

// Patterns represents a pattern configuration from YAML.
type Patterns struct {
	Name          string           `yaml:"name"`
	Regex         []string         `yaml:"regex"`
	CompiledRegex []*regexp.Regexp // Cached compiled regex
}

// Configs represents the YAML configuration structure.
type Configs struct {
	Patterns []Patterns `yaml:"patterns"`
}

// FinalResults represents the extracted data for a single match.
type FinalResults struct {
	Url     string
	Name    string
	Regex   string
	DataOut []string
}

var (
	cachedConfigs *Configs
	configOnce    sync.Once
	configError   error
)

// loadConfigs loads and parses the YAML configuration file, caching the result.
func loadConfigs() (*Configs, error) {
	configOnce.Do(func() {
		var configs Configs
		if err := yaml.Unmarshal(embeddedConfigs, &configs); err != nil {
			configError = fmt.Errorf("failed to unmarshal embedded config: %v", err)
			return
		}

		if len(configs.Patterns) == 0 {
			configError = fmt.Errorf("no patterns found in embedded config")
			return
		}

		for i, pattern := range configs.Patterns {
			configs.Patterns[i].CompiledRegex = make([]*regexp.Regexp, len(pattern.Regex))
			for j, regex := range pattern.Regex {
				compiled, err := regexp.Compile(regex)
				if err != nil {
					configError = fmt.Errorf("invalid regex %s: %v", regex, err)
					return
				}
				configs.Patterns[i].CompiledRegex[j] = compiled
			}
		}

		cachedConfigs = &configs
	})

	if configError != nil {
		return nil, configError
	}

	return cachedConfigs, nil
}

// ExtractData extracts data from the input string using regex patterns defined in configs.yaml.
// It returns a slice of FinalResult containing matched data or an error if the configuration cannot be loaded.
func ExtractData(data string) ([]FinalResults, error) {
	configs, err := loadConfigs()
	if err != nil {
		return nil, err
	}

	caser := cases.Title(language.English)
	results := make([]FinalResults, 0, len(configs.Patterns)*2) // Pre-allocate capacity

	for _, config := range configs.Patterns {
		for i, pattern := range config.Regex {
			matches := FilteredDataOutput(config.CompiledRegex[i].FindAllString(data, -1))
			if len(matches) > 0 {
				results = append(results, FinalResults{
					Name:    caser.String(config.Name),
					Regex:   pattern,
					DataOut: matches,
				})
			}
		}
	}

	return results, nil
}
