package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type FakJsBase struct {
	Args    string
	Targets []string
	Threads int
	Verbose bool
	*Client
}

type Results struct {
	Target  string
	RawData string
}

type FinalResults struct {
	Target  string
	Name    string
	Regex   string
	DataOut []string
}

func NewFakJs(target string, threads int, verbose bool) *FakJsBase {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var targets []string

	// Check if the -target flag is given
	if target != "" {
		if IsFile(target) {
			file, err := os.Open(target)
			if err != nil {
				fmt.Printf("%s: failed to open file: %v\n", ColoredText("red", "error"), err)
				os.Exit(1)
			}
			defer file.Close()

			lines, err := ReadLinesWithContext(ctx, file)
			if err != nil {
				fmt.Printf("%s: failed to read lines from file: %v\n", ColoredText("red", "error"), err)
				os.Exit(1)
			}

			targets = append(targets, lines...)
		} else {
			targets = append(targets, target)
		}
	} else {
		lines, err := ReadLinesWithContext(ctx, os.Stdin)
		if err != nil {
			fmt.Printf("%s: failed to read from stdin: %v\n", ColoredText("red", "error"), err)
		}
		targets = append(targets, lines...)
	}

	if len(targets) == 0 {
		fmt.Printf("%s: no targets provided\n", ColoredText("red", "error"))
		os.Exit(1)
	}

	client := NewClient()

	return &FakJsBase{
		Args:    target,
		Targets: targets,
		Threads: threads,
		Verbose: verbose,
		Client:  client,
	}
}

func (base FakJsBase) FakJsRun() {
	targets := make(chan string, base.Threads)
	results := make(chan Results, base.Threads)
	finalResults := make(chan FinalResults, base.Threads)

	var wgReq sync.WaitGroup

	for i := 0; i < base.Threads; i++ {
		wgReq.Add(1)

		go func() {
			defer wgReq.Done()

			for target := range targets {

				if strings.HasPrefix(target, "http") {
					resp, err := base.Client.Do("GET", target)
					if err != nil {
						FilteredVerboseOutput(base.Verbose, fmt.Sprintf("%s: fetching %s: %v", ColoredText("red", "error"), target, err))
						continue
					}

					body, err := io.ReadAll(resp.Body)
					if err != nil {
						FilteredVerboseOutput(base.Verbose, fmt.Sprintf("%s: reading response body for %s: %v", ColoredText("red", "error"), target, err))
						continue
					}

					results <- Results{
						Target:  target,
						RawData: string(body),
					}

					resp.Body.Close()

				} else {
					if base.Args == "" {
						results <- Results{
							Target:  "Unknown",
							RawData: target,
						}
					} else {
						results <- Results{
							Target:  base.Args,
							RawData: target,
						}
					}
				}
			}
		}()
	}

	// Read Targets from stdin
	go func() {
		for _, target := range base.Targets {
			targets <- target
		}
		close(targets)
	}()

	// Close results channel when all requests are done
	go func() {
		wgReq.Wait()
		close(results)
	}()

	var wgOut sync.WaitGroup

	wgOut.Add(1)
	go func() {
		for res := range results {
			data, err := ExtractData(res.RawData)
			if err != nil {
				FilteredVerboseOutput(base.Verbose, fmt.Sprintf("%s: extracting data for %s: %v", ColoredText("red", "error"), res.Target, err))
				continue
			}

			for _, out := range data {
				if len(out.DataOut) > 0 {
					if base.Verbose {
						fmt.Printf(
							"[%s] [%s] — {%s}\n%s\n",
							ColoredText("blue", out.Name),
							ColoredText("magenta", out.Regex),
							ColoredText("cyan", res.Target),
							ColoredText("green", strings.Join(out.DataOut, "\n")),
						)
					} else {
						fmt.Printf(
							"[%s] — {%s}\n%s\n",
							ColoredText("blue", out.Name),
							ColoredText("cyan", res.Target),
							ColoredText("green", strings.Join(out.DataOut, "\n")),
						)
					}

					finalResults <- FinalResults{
						Target:  res.Target,
						Name:    out.Name,
						Regex:   out.Regex,
						DataOut: out.DataOut,
					}
				}
			}
		}
		wgOut.Done()
	}()

	// Close finalResults channel when processing is done
	go func() {
		wgOut.Wait()
		close(finalResults)
	}()

	JsonReport(base.Verbose, finalResults)
}
