package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"unsafe"
)

var headerFlag = flag.String("h", "", "The header for the request")
var numRequestFlag = flag.Int("n", 1, "The number of requests")
var concurrentRequestFlag = flag.Int("c", 1, "The number of concurrent requests")
var requestTypeFlag = flag.String("m", "", "The type of the request")

func main() {
	flag.Parse()
	fmt.Println("Header: ", *headerFlag, "Address of Header Var: ", &headerFlag, unsafe.Sizeof(headerFlag))
	fmt.Println("Number of requests: ", *numRequestFlag, "Address of requests var: ", &numRequestFlag)
	fmt.Println("Concurrency Level: ", *concurrentRequestFlag, "Address of concurreny var: ", &concurrentRequestFlag)
	fmt.Println("Request Type: ", *requestTypeFlag, "Adddress of Request type var: ", &requestTypeFlag)

	//Since this returns an array we are going to have to make sure there is something in it before trying to access by index.
	fmt.Println("Web Server to bench: ", flag.Arg(0))
	createClient(flag.Arg(0))
}

func createClient(url string) {
	defer timer(time.Now(), "Http Request")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.Status
	fmt.Println(status)
}

func timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	log.Printf("%s took %s", name, timeTaken)
}
