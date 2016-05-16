package main

import (
	"flag"
	"net/http"
	"time"
	"log"
	"fmt"
)

var headerFlag = flag.String("h", "", "The header for the request")
var numRequestFlag = flag.Int("n", 1, "The number of requests")
var concurrentRequestFlag = flag.Int("c", 1, "The number of concurrent requests")
var requestTypeFlag = flag.String("m", "GET", "The type of the request")

type httpResponse struct {
	*http.Response
	err error
	timeTaken int64
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

func processResults(responseChannel chan httpResponse, numberOfRequests int) {
	for request := 0;  request < numberOfRequests; request++ {
		select {
			case response, ok := <-responseChannel:
				if ok {
					if response.err != nil {
						log.Println(response.err)
					} else {
						fmt.Println(time.Duration(response.timeTaken))
						if err := response.Body.Close(); err != nil {
							log.Println(response.err)
						}
					}
				}
		}
	}
}

func timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	log.Printf("%s took %s", name, timeTaken)
}

func main() {
	flag.Parse()
	webServerToBench := flag.Arg(0)
	httpRequestChannel := make(chan *http.Request)
	httpResponseChannel := make(chan httpResponse)
	go distributeWork(httpRequestChannel, *numRequestFlag, *requestTypeFlag, webServerToBench)
	defer timer(time.Now(), "Time of all requests")
	go workPool(httpRequestChannel, httpResponseChannel, *concurrentRequestFlag)
	processResults(httpResponseChannel, *numRequestFlag)
}
