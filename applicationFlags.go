package spotter

// ApplicationFlags is a struct representation of what important information is needed to run Spotter.
type ApplicationFlags struct {
	// headers will need to change to map.
	headers            string
	requestNumber      int
	concurrentRequests int
	requestType        string
	usage              func()
}

// NewApplicationFlags creates a new instance of our application flags that can be mapped to.
func NewApplicationFlags(headers string, requestNumber int, concurrentRequestNumber int, requestType string) *ApplicationFlags {
	return &ApplicationFlags{headers: headers, requestNumber: requestNumber, concurrentRequests: concurrentRequestNumber, requestType: requestType}
}

// SetHeaders will set the header value for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetHeaders(headers string) {
	a.headers = headers
}

// SetRequestNumber will set the request number for the associated instance ApplicationFlags.
func (a *ApplicationFlags) SetRequestNumber(requestNumber int) {
	a.requestNumber = requestNumber
}

// SetConcurrentRequestNumber will set the the level of concurrency for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetConcurrentRequestNumber(concurrentRequestNumber int) {
	a.concurrentRequests = concurrentRequestNumber
}

// SetRequestType will set the type of reuqest to make for this associated instance of ApplicationFlags.
func (a *ApplicationFlags) SetRequestType(requestType string) {
	a.requestType = requestType
}

// GetHeaders is an accessor to the headers property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetHeaders() string {
	return a.headers
}

// GetRequestNumber is an accessor to the request number property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetRequestNumber() int {
	return a.requestNumber
}

// GetConcurrentRequestNumber is an accessor to the concurrency level property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetConcurrentRequestNumber() int {
	return a.concurrentRequests
}

// GetRequestType is an accessor to the request type property for the associated instance of ApplicationFlags.
func (a *ApplicationFlags) GetRequestType() string {
	return a.requestType
}

// SetUsage accepts a callable to be invoked when the usage of ApplicationFlags is desired.
func (a *ApplicationFlags) SetUsage(b func()) {
	a.usage = b
}

// GetUsage is an accessor to the usage callable to be invoked during a fail state.
func (a *ApplicationFlags) GetUsage() func() {
	return a.usage
}
