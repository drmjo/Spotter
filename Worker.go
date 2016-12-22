package main

import (
	"log"
	"time"
	"net/http"
)

// Dial: (&net.Dialer {
// 				Timeout: 30 * time.Second,
// 				KeepAlive: 30 * time.Second,
// 		}).Dial,

// NOTE: Can configure SSL and redirect policy here later.
var transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

var client = &http.Client{
	Transport: transport,
}

type Worker struct {
	ID         int
	WorkerPool chan chan WorkRequest
	JobChannel chan WorkRequest
	quit       chan bool
}

func NewWorker(id int, workerPool chan chan WorkRequest) Worker {
	worker := Worker{
		ID:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan WorkRequest),
		quit:       make(chan bool)}

	return worker
}

func (w *Worker) Start(resultChannel chan WorkResponse) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//fmt.Printf("Worker %d received a request\n", w.ID)

				start := time.Now()
				resp, err := client.Do(job.HTTPRequest)
				end := time.Now()
				
				if err != nil {
					log.Fatalf("Couldn't complete HTTP Request %v", err)
				}
				response := NewWorkResponse(resp, start, end)

				resultChannel <- *response
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
