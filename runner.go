package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Results struct {
	StatusCode int
	Url        string
	Body       string
}

func FakJsRunner(concurrency int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	std := bufio.NewScanner(os.Stdin)
	if !std.Scan() && std.Err() == nil {
		fmt.Printf("%s: No input provided\n", ColoredText("red", "error"))
		return nil
	}

	targets := make(chan string, concurrency)
	results := make(chan Results, concurrency)
	finalResults := make(chan FinalResults, concurrency)

	client := NewClient()

	var wgReq sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wgReq.Add(1)

		go func() {
			defer wgReq.Done()

			for url := range targets {

				resp, err := client.Do("GET", url)
				if err != nil {
					fmt.Printf("%s: fetching %s: %v", ColoredText("red", "error"), url, err)
					continue
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("%s: reading response body for %s: %v\n", ColoredText("red", "error"), url, err)
					continue
				}

				results <- Results{
					StatusCode: resp.StatusCode,
					Url: url,
					Body: string(body),
				}
			}
		}()
	}

	// Read URLs from stdin
	go func() {
		defer close(targets)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line := strings.TrimSpace(std.Text())
				if line != "" {
					targets <- line
				}
				if !std.Scan() {
					return
				}
			}
		}
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
            data, err := ExtractData(res.Body)
            if err != nil {
                fmt.Printf("%s: extracting data for %s: %v\n", ColoredText("red", "error"), res.Url, err)
                continue
            }
            for _, out := range data {
                if len(out.DataOut) > 0 {
                    fmt.Printf(
                        "[%s] [%s] --- [%s] --- {%s}\n",
                        ColoredText("blue", out.Name),
                        ColoredText("magenta", out.Regex),
                        ColoredText("green", strings.Join(out.DataOut, ", ")),
                        ColoredText("cyan", res.Url),
                    )
                    finalResults <- FinalResults{
						Url: res.Url,
						Name: out.Name,
						Regex: out.Regex,
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

	return JsonReport(finalResults)
}
