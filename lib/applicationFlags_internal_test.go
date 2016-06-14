package lib

import(
	"testing"
)

func TestSetAndGetHeaders(t *testing.T) {
	fakeHeaderValues := "some random header values"
	testApplicationFlags := &ApplicationFlags{}
	testApplicationFlags.SetHeaders(fakeHeaderValues)

	if headers := testApplicationFlags.GetHeaders(); fakeHeaderValues != headers {
		t.Errorf("Error Testing Set/Get Headers. Received %v but was expecting %v", headers, fakeHeaderValues)
	}
}

func TestSetAndGetRequestNumber(t *testing.T) {
	fakeRequestNumber := 10
	testApplicationFlags := &ApplicationFlags{}
	testApplicationFlags.SetRequestNumber(fakeRequestNumber)

	if requests := testApplicationFlags.GetRequestNumber(); fakeRequestNumber != requests {
		t.Errorf("Error Testing Set/Get Request Number. Received %v but was expecting %v", requests, fakeRequestNumber)
	}
}

func TestSetAndGetConcurrentRequestNumber(t *testing.T) {
	fakeConcurrencyLevel := 5 
	testApplicationFlags := &ApplicationFlags{}
	testApplicationFlags.SetConcurrentRequestNumber(fakeConcurrencyLevel)

	if cLevel := testApplicationFlags.GetConcurrentRequestNumber(); cLevel != fakeConcurrencyLevel {
		t.Errorf("Error Testing Set/Get Concurrent Request Number. Received %v but was expecting %v", cLevel, fakeConcurrencyLevel)
	}
}

func TestSetAndGetRequestType(t *testing.T) {
	fakeRequestType := "GET"
	testApplicationFlags := &ApplicationFlags{}
	testApplicationFlags.SetRequestType(fakeRequestType)

	if rType := testApplicationFlags.GetRequestType(); rType != fakeRequestType {
		t.Errorf("Error Testing Set/Get Request Type. Received %v but was expecting %v", rType, fakeRequestType)
	}
}

func TestSetAndGetURL(t *testing.T) {
	fakeURL := "example.com/elves"
	testApplicationFlags := &ApplicationFlags{}
	testApplicationFlags.SetURL(fakeURL)

	if url := testApplicationFlags.GetURL(); url != fakeURL {
		t.Errorf("Error testing Set/Get URL. Received %v but was expecting %v", url, fakeURL)
	}
}