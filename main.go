package main

import (
	"flag"
	"net/http"
	"time"
	"log"
	"fmt"
	"os"
)

var requestFlags cliFlags

func init() {
	 requestFlags.header = flag.String("h", "", "The header for the request")
	 requestFlags.requestNumber = flag.Int("n", 1, "The number of requests")
	 requestFlags.concurrentRequests = flag.Int("c", 1, "The number of concurrent requests")
	 requestFlags.requestType = flag.String("r", "GET", "The type of the request")
}

// Probably going to throw this out.
func timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	fmt.Println(name, timeTaken)
}

func main() {
	flag.Parse()
	webServerToBench := flag.Arg(0)
	httpRequestChannel := make(chan *http.Request)
	httpResponseChannel := make(chan httpResponse)
    defer timer(time.Now(), "Time to Complete: \t")
	go distributeWork(httpRequestChannel, *requestFlags.requestNumber, *requestFlags.requestType, webServerToBench)
	go workPool(httpRequestChannel, httpResponseChannel, *requestFlags.concurrentRequests)
	successfulConnections, totalTime := processResults(httpResponseChannel, *requestFlags.requestNumber)
	if successfulConnections == 0 {
		log.Println("Zero successful connections. Exiting.")
		os.Exit(1)
	}
	averageTime := time.Duration(totalTime/successfulConnections)
	fmt.Println("Web Server To Bench: \t", webServerToBench)
	fmt.Println("Number of Requests \t", *requestFlags.requestNumber)
	fmt.Println("Successful Requests: \t", successfulConnections)
	fmt.Println("Concurrency Level: \t", *requestFlags.concurrentRequests)
	fmt.Println("Total Request Time: \t", time.Duration(totalTime))
	fmt.Println("Average Request Time: \t", averageTime)
}
