package main

import (
	"net/http"
)

type WorkRequest struct {
	NumberOfWorkers  int
	NumberOfRequests int
	HTTPRequest      *http.Request
}

func NewWorkRequest(workerNum int, reqNumber int, httpReq *http.Request) WorkRequest {
	workRequest := WorkRequest{
		NumberOfWorkers:  workerNum,
		NumberOfRequests: reqNumber,
		HTTPRequest:      httpReq,
	}

	return workRequest
}
