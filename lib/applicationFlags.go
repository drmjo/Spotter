package lib

// ApplicationFlags is a struct representation of what important information is needed to run Spotter.
type ApplicationFlags struct {
	// headers will need to change to map.
	headers            string
	requestNumber      int
	concurrentRequests int
	requestType        string
}

// NewApplicationFlags creates a new instance of our application flags that can be mapped to.
func NewApplicationFlags(headers string, requestNumber int, concurrentRequestNumber int, requestType string) *ApplicationFlags {
	return &ApplicationFlags{headers: headers, requestNumber: requestNumber, concurrentRequests: concurrentRequestNumber, requestType: requestType}
}

// SetHeaders will set the header value for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetHeaders(headers string) {
	a.headers = headers
}

// SetRequestNumber will set the request number for the associated instance ApplicationFlags.
func (a *ApplicationFlags) SetRequestNumber(requestNumber int) {
	a.requestNumber = requestNumber
}

// SetConcurrentRequestNumber will set the the level of concurrency for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetConcurrentRequestNumber(concurrentRequestNumber int) {
	a.concurrentRequests = concurrentRequestNumber
}

// SetRequestType will set the type of reuqest to make for this associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetRequestType(requestType string) {
	a.requestType = requestType
}

// GetHeaders is an accessor to the headers property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetHeaders() string {
	return a.headers
}

// GetRequestNumber is an accessor to the request number property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetRequestNumber() int {
	return a.requestNumber
}

// GetConcurrentRequestNumber is an accessor to the concurrency level property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetConcurrentRequestNumber() int {
	return a.concurrentRequests
}

// GetRequestType is an accessor to the request type property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetRequestType() string {
	return a.requestType
}

// import (
// 	"net/http"
// 	"log"
// 	"time"
// )

// func distributeWork(requestChannel chan *http.Request, numberOfRequests int, requestType string, url string) {
// 	for i := 0; i < numberOfRequests; i++ {
// 		request, err := http.NewRequest(requestType, url, nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		requestChannel <- request
// 	}
// }

// func workPool(requestChannel chan *http.Request, responseChannel chan httpResponse, concurrencyLevel int) {
// 	transport := &http.Transport{}
// 	for i := 0; i < concurrencyLevel; i++ {
// 		go worker(transport, requestChannel, responseChannel)
// 	}
// }

// func worker(transport *http.Transport, requestChannel chan *http.Request, responseChannel chan httpResponse) {
// 	for work := range requestChannel {
// 		start := time.Now()
// 		responseRoundtrip, err := transport.RoundTrip(work)
// 		timeTaken := time.Since(start)
// 		response := httpResponse{responseRoundtrip, err, int64(timeTaken)}
// 		responseChannel <- response
// 	}
// }

// func processResults(responseChannel chan httpResponse, numberOfRequests int) (int64, int64) {
// 	var successfulConnections, totalTime int64
// 	for request := 0; request < numberOfRequests; request++ {
// 		select {
// 			case response, ok := <-responseChannel:
// 				if ok {
// 					if response.err != nil {
// 						log.Println(response.err)
// 					} else {
// 						successfulConnections++
// 						totalTime += response.timeTaken
// 						if err := response.Body.Close(); err != nil {
// 							log.Println(response.err)
// 						}
// 					}
// 				}
// 		}
// 	}
// 	return successfulConnections, totalTime
// }
