package main

import (
	"net/http"
	"log"
	"time"
)

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
