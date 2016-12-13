package main

import (
	"flag"
	"fmt"
)

// TODO: Investigate moving all of this to the cobra command line package.
var (
	NumberWorkers  = flag.Int("c", 1, "The number of workers.")
	NumberRequests = flag.Int("n", 1, "The number of requests.")
	RequestType    = flag.String("r", "GET", "The request type.")
	//  NOTE:look up the json comment annotation on structs and decoding strings onto structs.
	// RequestData = flag.String("d", "", "The request data.")
	// RequestHeaders = flag.String("h", "", "The request headers")
)

func main() {
	flag.Parse()
	// , *RequestData, *RequestHeaders
	work := NewWorkRequest(*NumberWorkers, *NumberRequests, *RequestType)

	fmt.Println("Let's bench, brah!")
	dispatcher := NewDispatcher(&work)
	dispatcher.Run()
}
