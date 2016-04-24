package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	"os"
)

var headerFlag = flag.String("h", "", "The header for the request")
var numRequestFlag = flag.Int("n", 1, "The number of requests")
var concurrentRequestFlag = flag.Int("c", 1, "The number of concurrent requests")
var requestTypeFlag = flag.String("m", "GET", "The type of the request")

func main() {
	flag.Parse()
	fmt.Println("Header: ", *headerFlag, "Address of Header Var: ", &headerFlag)
	fmt.Println("Number of requests: ", *numRequestFlag, "Address of requests var: ", &numRequestFlag)
	fmt.Println("Concurrency Level: ", *concurrentRequestFlag, "Address of concurreny var: ", &concurrentRequestFlag)
	fmt.Println("Request Type: ", *requestTypeFlag, "Adddress of Request type var: ", &requestTypeFlag)
	fmt.Println("Web Server to bench: ", flag.Arg(0))
	sendRequest(flag.Arg(0), *numRequestFlag, *requestTypeFlag)
}

func sendRequest(url string, requestNumber int, requestType string) {
	client := &http.Client{}
	timeList := []int64{}
	determineRequestType(&requestType)
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer timer(time.Now(), "Request time plus setup")
	for requestNumber > 0 {
		requestTime := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		status := resp.Status
		testTime := time.Since(requestTime)
		timeList = append(timeList, int64(testTime))
		fmt.Printf("%s \tRequest time: %s \n", status, testTime)
		requestNumber--
	}
	fmt.Printf("Average request time was: %s \n", determineAverageRequestTime(timeList))
}

func determineAverageRequestTime(timeList []int64) time.Duration {
	var averageTime int64
	for i := range timeList {
		averageTime += timeList[i]
	}
	return time.Duration(averageTime/int64(len(timeList)))
}

func determineRequestType(requestType *string) {
	switch strings.ToUpper(*requestType) {
	case "GET":
		*requestType = "GET"
		return
	case "POST":
		*requestType = "POST"
		return
	case "PUT":
		*requestType = "PUT"
		return
	case "DELETE":
		*requestType = "DELETE"
		return
	}
	log.Fatalf("Cannot process the request type of %s", *requestType)
	os.Exit(1)
}

func timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	log.Printf("%s took %s", name, timeTaken)
}
