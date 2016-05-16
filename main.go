package main

import (
	"flag"
	"fmt"
	"http"
)

var headerFlag = flag.String("h", "", "The header for the request")
var numRequestFlag = flag.Int("n", 1, "The number of requests")
var concurrentRequestFlag = flag.Int("c", 1, "The number of concurrent requests")
var requestTypeFlag = flag.String("m", "GET", "The type of the request")

func main() {
	flag.Parse()
	webServerToBench := flag.Arg(0)
	httpRequestChannel := make(chan *http.Request)
	httpResponseChannel := make(chan HttpResponse)
	timeList := []int64{}
}
