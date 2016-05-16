package main 

import(
	"net/http"
	"log"
	"fmt"
	"time"
)

// TODO this will likely hav to receive from a channel of struct array for url and request type.
func Worker(id int, requestChannel chan *http.Request, responseChannel chan Response) {
	for work := range jobs {
		fmt.Println("worker", id, "request numer", requestNumber)
		sendRequest(work,  )
	}
}

// TODO this will likely have request number removed from the signature.
func sendRequest(url string, requestNumber int, requestType string) {
	client := &http.Client{}
	determineRequestType(&requestType)
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for requestNumber > 0 {
		requestTime := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		status := resp.Status
		testTime := time.Since(requestTime)
		timeList = append(timeList, int64(testTime))
		fmt.Printf("%s \tRequest time: %s \n", status, testTime)
		requestNumber--
	}
	fmt.Printf("Average request time was: %s \n", determineAverageRequestTime(timeList))
}


//EVERYTHING BENEATH HERE CAN MOST LIKELY STAY.

func determineAverageRequestTime(timeList []int64) time.Duration {
	var averageTime int64
	for i := range timeList {
		averageTime += timeList[i]
	}
	return time.Duration(averageTime/int64(len(timeList)))
}

func determineRequestType(requestType *string) {
	switch strings.ToUpper(*requestType) {
	case "GET":
		*requestType = "GET"
	case "POST":
		*requestType = "POST"
	case "PUT":
		*requestType = "PUT"
	case "DELETE":
		*requestType = "DELETE"
	default:
		log.Fatalf("Cannot process the request type of %s", *requestType)
		os.Exit(1)
	}
}

//defer timer(time.Now(), "Request time plus setup")
func timer(start time.Time, name string) {
	timeTaken := time.Since(start)
	log.Printf("%s took %s", name, timeTaken)
}
