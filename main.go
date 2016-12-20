package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type headers []string

var (
	workers        int
	requests       int
	requestMethod  string
	requestBody    string
	requestHeaders headers
)

func init() {
	flag.IntVar(&workers, "c", 1, "The number of workers.")
	flag.IntVar(&requests, "n", 1, "The number of requests.")
	flag.StringVar(&requestMethod, "r", "GET", "The request type.")
	flag.StringVar(&requestBody, "d", "", "The request data.")
	flag.Var(&requestHeaders, "h", "The request headers")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	urlDirty := args[0]
	urlClean := checkURL(urlDirty)
	httpRequest := createHttpRequest(requestMethod, requestBody, requestHeaders, urlClean)
	work := NewWorkRequest(workers, requests, httpRequest)

	fmt.Println("Starting Benchmarks...")
	dispatcher := NewDispatcher(&work)
	dispatcher.Run()
}

func checkURL(uri string) *url.URL {
	url, error := url.Parse(uri)
	if error != nil {
		log.Fatalf("Could not parse url %q: %v", url, error)
	}

	return url
}

func createHttpRequest(requestMethod string, requestBody string, requestHeaders headers, url *url.URL) *http.Request {
	req, err := http.NewRequest(requestMethod, url.String(), makeHttpBody(requestBody))
	if err != nil {
		log.Fatalf("Can't create HTTP Request %v", err)
	}

	for _, value := range requestHeaders {
		key, value := extractHeaderValue(value)
		req.Header.Add(key, value)
	}

	return req
}

func extractHeaderValue(header string) (string, string) {
	splitHeader := strings.Split(header, ":")
	if len(splitHeader) != 2 {
		log.Fatalf("Malformed Request Header %v", header)
	}

	return splitHeader[0], splitHeader[1]
}

func makeHttpBody(body string) io.Reader {
	return strings.NewReader(body)
}

func (h headers) String() string {
	var buffer bytes.Buffer
	for _, value := range h {
		buffer.WriteString(value)
	}

	return buffer.String()
}

func (h *headers) Set(value string) error {
	*h = append(*h, value)
	return nil
}
