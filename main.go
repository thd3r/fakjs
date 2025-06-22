package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var target string
	flag.StringVar(&target, "target", "", " single target or file containing multiple targets")

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "show verbose output")

	var version bool
	flag.BoolVar(&version, "version", false, "show fakjs version")

	var concurrency int
	flag.IntVar(&concurrency, "threads", 40, "number of concurrent threads")

	flag.Parse()

	if version {
		fmt.Printf("Fakjs current version: %s\n", strings.ReplaceAll(CurrentVersion, "v", ""))
		os.Exit(0)
	}

	var banner = fmt.Sprintf(`
	███████╗ █████╗ ██╗  ██╗     ██╗███████╗
	██╔════╝██╔══██╗██║ ██╔╝     ██║██╔════╝
	█████╗  ███████║█████╔╝      ██║███████╗
	██╔══╝  ██╔══██║██╔═██╗ ██   ██║╚════██║
	██║     ██║  ██║██║  ██╗╚█████╔╝███████║
	╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝ ╚════╝ ╚══════╝
			  %s																		
	`, Version())

	fmt.Println(banner)

	fmt.Println(":: Fakjs — extract sensitive info from JS")
	fmt.Printf(":: Generating report at %s\n", FilePath)

	runner := NewFakJs(target, concurrency, verbose)
	runner.FakJsRun()
}
