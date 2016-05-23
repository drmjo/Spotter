package main

import "log"

func processResults(responseChannel chan httpResponse, numberOfRequests int) (int64, int64) {
	var successfulConnections, totalTime int64
	for request := 0; request < numberOfRequests; request++ {
		select {
			case response, ok := <-responseChannel:
				if ok {
					if response.err != nil {
						log.Println(response.err)
					} else {
						successfulConnections++
						totalTime += response.timeTaken
						if err := response.Body.Close(); err != nil {
							log.Println(response.err)
						}
					}
				}
		}
	}
	return successfulConnections, totalTime
}
