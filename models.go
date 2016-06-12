package main 

import "net/http"



type httpResponse struct {
	*http.Response
	err error
	timeTaken int64
}
