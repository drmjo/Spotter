package lib 

type MockApplicationFlags struct {
	headers string
	requestNumber int
	concurrentRequests int
	requestType string	
	url string
}

func (m *MockApplicationFlags) SetHeaders(headers string) {
	m.headers = headers
}

func (m *MockApplicationFlags) SetRequestNumber(requestNumber int) {
	m.requestNumber = requestNumber
}

func (m *MockApplicationFlags) SetConcurrentRequestNumber(cLevel int){
	m.concurrentRequests = cLevel
}

func (m *MockApplicationFlags) SetRequestType(rType string) {
	m.requestType = rType
}

func (m *MockApplicationFlags) SetURL(url string) {
	m.url = url
}

func (m *MockApplicationFlags) GetHeaders() string {
	return m.headers
}

func (m *MockApplicationFlags) GetRequestNumber() int {
	return m.requestNumber
}

func (m *MockApplicationFlags) GetConcurrentRequestNumber() int {
	return m.concurrentRequests
}

func (m *MockApplicationFlags) GetRequestType() string {
	return m.requestType
}

func (m *MockApplicationFlags) GetURL() string {
	return m.url
}

// Neat little trick from Nanobase. https://github.com/codykrainock/nanobase/blob/master/mock_collection_internal_test.go
var _ flagsInterface = (*MockApplicationFlags)(nil)







