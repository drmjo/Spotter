package main 

import(
	"net/http"
	"log"
	"fmt"
	"time"
)

func DistributeWork(requestChannel chan *http.Request, numberOfRequests int, requestType string, url string) {
	for i := 0; i < numberOfRequests; i++ {
		request, err := http.NewRequest(requestType, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		requestChannel <- request
	}
}

func WorkPool(requestChannel chan *http.Request, responseChannel chan HttpResponse, concurrencyLevel int) {
	transport := &http.transport{}
	for i := 0; i < concurrencyLevel; i++ {
		go worker(t, requestChannel, responseChannel)
	}
}

func worker(transport *http.Transport, requestChannel chan *http.Request, responseChannel chan HttpResponse) {
	for work := range requestChannel {
		start := time.Now()
		responseRoundtrip, err := transport.RoundTrip(work)
		timeTaken := time.Since(start)
		response := HttpResponse{responseRoundtrip, err, int64(timeTaken)}
		responseChannel <- response
	}
}

func ProcessResults(responseChannel chan Response, numberOfRequests int) {
	for request := 0;  request < numberOfRequests; request++ {
		select {
			case response, ok := <-responseChannel:
				if ok != nil {
					if response.err != nil {
						log.Println(response.err)
					} else {
						if err := response.Body.Close(); err != nil {
							log.Println(response.err)
						}
					}
				}
		}
	}
}

//defer timer(time.Now(), "Request time plus setup")
func Timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	log.Printf("%s took %s", name, timeTaken)
}
