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
	`, VERSION)

	fmt.Println(banner)
}

func main() {
	var concurrency int
	flag.IntVar(&concurrency, "c", 40, "number of concurrent goroutines")

	flag.Parse()

	fmt.Println(":: Fakjs — extract sensitive info from JS")
	fmt.Printf(":: Generating report at %s\n", FilePath)

	if err := FakJsRunner(concurrency); err != nil {
		fmt.Printf("%s: %v", ColoredText("red", "error"), err)
	}
}
