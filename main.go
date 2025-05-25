package main

import (
	"flag"
	"fmt"
)

func init() {
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
}

func main() {
	var target string
	flag.StringVar(&target, "target", "", " single target or file containing multiple targets")

	var concurrency int
	flag.IntVar(&concurrency, "threads", 40, "number of concurrent threads")

	flag.Parse()

	fmt.Println(":: Fakjs — extract sensitive info from JS")
	fmt.Printf(":: Generating report at %s\n", FilePath)

	runner := NewFakJs(target, concurrency)
	runner.FakJsRun()
}
