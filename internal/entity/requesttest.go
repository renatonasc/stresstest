package entity

type RequestTest struct {
	Request                    Request
	NumberOfRequests           int
	NumberOfConcurrentRequests int
	NumberOfRequestsPerSecond  int
}
