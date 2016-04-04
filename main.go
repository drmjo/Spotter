package main

import (
	"flag"
	"fmt"
)
// Entry point to our dope program.
func main() {


	// TODO: Think about what the default header for this should be
	headerFlag := flag.String("h", "", "The header for the request")

	// Just hit the endpoint once
	numRequestFlag := flag.Int("n", 1, "The number of requests")

	// Probably concurrency level 1 unless otherwise specified
	concurrentRequestFlag := flag.Int("c", 1, "The number of concurrent requests")

	// TODO: Think about default request type. GET?
	requestTypeFlag := flag.String("m", "", "The type of the request")

	// Parse the command line flags
	flag.Parse();

	// Make sure we are getting what we 
	fmt.Println("Header: ", *headerFlag)
	fmt.Println("Number of requests: ", *numRequestFlag)
	fmt.Println("Concurrency Level: ", *concurrentRequestFlag)
	fmt.Println("Request Type: ", *requestTypeFlag)
	// Since this returns an array we are going to have to make sure there is something in it before trying to access by index.
	fmt.Println("Web Server to bench: ", flag.Args())
}

