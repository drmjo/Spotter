package main

import (
	"flag"
	"fmt"
	"os"

	spotter "github.com/hunterel/spotter/lib"
)

func main() {
	cliFlags := spotter.ApplicationFlags{}
	headers := flag.String("h", "", "Headers for the http request")
	requests := flag.Int("n", 1, "Number of http requests to make")
	concurrency := flag.Int("c", 1, "Concurrency level for making http requests")
	requestType := flag.String("r", "GET", "Type of http request to make")
	usage := func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("  spotter -r 10 -c 2 -r GET https://example.com\n")
		flag.PrintDefaults()

	}

	flag.Parse()
	if flag.NArg() == 0 {
		usage()
		os.Exit(1)
	}
	cliFlags.SetHeaders(*headers)
	cliFlags.SetRequestNumber(*requests)
	cliFlags.SetRequestType(*requestType)
	cliFlags.SetConcurrentRequestNumber(*concurrency)
	cliFlags.SetURL(flag.Arg(0))

	fmt.Printf("%+v", cliFlags)
	// httpRequestChannel := make(chan *http.Request)
	// httpResponseChannel := make(chan httpResponse)
	// go distributeWork(httpRequestChannel, *cliFlags.requestNumber, *cliFlags.requestType, webServerToBench)
	// go workPool(httpRequestChannel, httpResponseChannel, *cliFlags.concurrentRequests)
	// successfulConnections, totalTime := processResults(httpResponseChannel, *cliFlags.requestNumber)
	// if successfulConnections == 0 {
	// 	log.Println("Zero successful connections. Exiting.")
	// 	os.Exit(1)
	// }
	// averageTime := time.Duration(totalTime / successfulConnections)
	// fmt.Println("Web Server To Bench: \t", webServerToBench)
	// fmt.Println("Number of Requests \t", *cliFlags.requestNumber)
	// fmt.Println("Successful Requests: \t", successfulConnections)
	// fmt.Println("Concurrency Level: \t", *cliFlags.concurrentRequests)
	// fmt.Println("Total Request Time: \t", time.Duration(totalTime))
	// fmt.Println("Average Request Time: \t", averageTime)
}
