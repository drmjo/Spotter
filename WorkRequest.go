package main

type WorkRequest struct {
	NumberOfWorkers  int
	NumberOfRequests int
	RequestVerb      string
}

func NewWorkRequest(workerNum int, reqNumber int, verb string) WorkRequest {
	workRequest := WorkRequest{
		NumberOfWorkers:  workerNum,
		NumberOfRequests: reqNumber,
		RequestVerb:      verb,
	}

	return workRequest
}
