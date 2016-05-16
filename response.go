package main 

import "http"

type HttpResponse struct (
	*http.Response
	err error
)
