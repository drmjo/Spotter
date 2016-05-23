package main 

import "net/http"

type cliFlags struct {
	header *string
	requestNumber *int
	concurrentRequests *int
	requestType *string
}

type httpResponse struct {
	*http.Response
	err error
	timeTaken int64
}
