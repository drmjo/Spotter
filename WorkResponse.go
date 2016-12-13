package main

type WorkResponse struct {
	// TODO: Fill in with pertinent response details
	message string
	id      int
}

func NewWorkResponse() *WorkResponse {
	workResponse := &WorkResponse{}
	return workResponse
}
