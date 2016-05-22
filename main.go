package main

import (
	"flag"
	"net/http"
	"time"
	"log"
	"fmt"
)

var requestFlags cliFlags

func init() {
	 requestFlags.header = flag.String("h", "", "The header for the request")
	 requestFlags.requestNumber = flag.Int("n", 1, "The number of requests")
	 requestFlags.concurrentRequests = flag.Int("c", 1, "The number of concurrent requests")
	 requestFlags.requestType = flag.String("r", "GET", "The type of the request")
}

func distributeWork(requestChannel chan *http.Request, numberOfRequests int, requestType string, url string) {
	for i := 0; i < numberOfRequests; i++ {
		request, err := http.NewRequest(requestType, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		requestChannel <- request
	}
}

func workPool(requestChannel chan *http.Request, responseChannel chan httpResponse, concurrencyLevel int) {
	transport := &http.Transport{}
	for i := 0; i < concurrencyLevel; i++ {
		go worker(transport, requestChannel, responseChannel)
	}
}

func worker(transport *http.Transport, requestChannel chan *http.Request, responseChannel chan httpResponse) {
	for work := range requestChannel {
		start := time.Now()
		responseRoundtrip, err := transport.RoundTrip(work)
		timeTaken := time.Since(start)
		response := httpResponse{responseRoundtrip, err, int64(timeTaken)}
		responseChannel <- response
	}
}

func processResults(responseChannel chan httpResponse, numberOfRequests int) (int64, int64) {
	var successfulConnections, totalTime int64
	for request := 0;  request < numberOfRequests; request++ {
		select {
			case response, ok := <-responseChannel:
				if ok {
					if response.err != nil {
						log.Println(response.err)
					} else {
						successfulConnections++
						totalTime += response.timeTaken
						if err := response.Body.Close(); err != nil {
							log.Println(response.err)
						}
					}
				}
		}
	}
	return successfulConnections, totalTime
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
	averageTime := time.Duration(totalTime/successfulConnections)
	fmt.Println("Web Server To Bench: \t", webServerToBench)
	fmt.Println("Number of Requests \t", *requestFlags.requestNumber)
	fmt.Println("Successful Requests: \t", successfulConnections)
	fmt.Println("Concurrency Level: \t", *requestFlags.concurrentRequests)
	fmt.Println("Total Request Time: \t", time.Duration(totalTime))
	fmt.Println("Average Request Time: \t", averageTime)
}
