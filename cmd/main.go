package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hunterel/spotter"
)

func main() {
	cliFlags := spotter.NewApplicationFlags()
	cliFlags.set = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("  spotter -r 10 -c 2 -r GET https://example.com\n")
		flag.PrintDefaults()
	}

	cliFlags := applicationFlags{}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cliFlags)

	// webServerToBench := flag.Arg(0)
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
