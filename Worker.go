package main

import (
	"fmt"
	"log"
	"time"
)

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
				fmt.Printf("Worker %d received a request\n", w.ID)

				start := time.Now()
				resp, err := client.Do(job.HTTPRequest)
				if err != nil {
					log.Fatalf("Couldn't complete HTTP Request %v", err)
				}
				end := time.Now()
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
