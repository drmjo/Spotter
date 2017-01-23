package main

import (
	"fmt"
	"time"
	"io/ioutil"
)

type Dispatcher struct {
	WorkerPool chan chan WorkRequest
	work       WorkRequest
}

var ResultChannel chan WorkResponse

func NewDispatcher(work *WorkRequest) *Dispatcher {
	pool := make(chan chan WorkRequest, (*work).NumberOfWorkers)
	return &Dispatcher{WorkerPool: pool, work: *work}
}

func (d *Dispatcher) Run() {

	ResultChannel := make(chan WorkResponse, d.work.NumberOfRequests)

	// NOTE: start goroutines equal to worker number then go through request count and handout requests.
	for i := 0; i < d.work.NumberOfWorkers; i++ {
		//fmt.Printf("Starting Worker: %d \n", i+1)
		worker := NewWorker(i, d.WorkerPool)
		worker.Start(ResultChannel)
	}

	go d.dispatch()

	// NOTE: Will only store the result time per status code to start.
	aggregateMap := make(map[int][]time.Duration)

	for i := 0; i < d.work.NumberOfRequests; i++ {
		result := <-ResultChannel
		display := &statsDisplay{}
		bodyBytes, err := ioutil.ReadAll(result.HTTPResponse.Body)
		var bodyString string
		if err != nil {
			bodyString = "Could not read request body"
		} else {
			bodyString = string(bodyBytes)
		}
		
		aggregateMap[result.HTTPResponse.StatusCode] =
		//fmt.Printf("Status code %d, time %s\n", result.HTTPResponse.StatusCode, result.End.Sub(result.Start))
		aggregateMap[result.HTTPResponse.StatusCode] = append(aggregateMap[result.HTTPResponse.StatusCode], result.End.Sub(result.Start))
	}

	resultMap := make(map[int]float64)

	for k, v := range aggregateMap {
		var averageTime float64
		if len(v) != 0 {
			var totalTime int64
			for _, i := range v {
				totalTime += int64(i)
			}
			averageTimeNS := (float64(totalTime) / float64(len(v)))
			averageTime = (averageTimeNS / float64(1000)) / float64(1000)
		}

		resultMap[k] = averageTime
	}

	for k, v := range resultMap {
		fmt.Printf("Average time for status code %d: %fms\n", k, v)
	}
}

type statsDisplay struct {
	timeTake time.Duration
	response string
}

func (d *Dispatcher) dispatch() {
	for i := 0; i < d.work.NumberOfRequests; i++ {
		go func(work WorkRequest) {
			jobChannel := <-d.WorkerPool

			jobChannel <- d.work
		}(d.work)
	}
}
