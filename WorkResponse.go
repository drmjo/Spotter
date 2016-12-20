package main

import (
	"net/http"
	"time"
)

type WorkResponse struct {
	HTTPResponse *http.Response
	Start        time.Time
	End          time.Time
}

func NewWorkResponse(httpResp *http.Response, start time.Time, end time.Time) *WorkResponse {
	workResponse := &WorkResponse{
		HTTPResponse: httpResp,
		Start:        start,
		End:          end,
	}
	return workResponse
}
