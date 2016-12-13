package main

import (
	"fmt"
)

type Worker struct {
	ID         int
	WorkerPool chan chan string
	JobChannel chan string
	quit       chan bool
}

func NewWorker(id int, workerPool chan chan string) Worker {
	worker := Worker{
		ID:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan string),
		quit:       make(chan bool)}

	return worker
}

func (w *Worker) Start(resultChannel chan WorkResponse) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Worker %d received a request of type %s\n", w.ID, job)
				response := NewWorkResponse()
				response.message = "Processed Request"
				response.id = w.ID
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
