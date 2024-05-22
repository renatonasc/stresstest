package entity

type RequestTestResult struct {
	RequestTest               RequestTest
	NumberOfCompletedRequests int
	NumberOfFailedRequests    int
	TimeMeanByRequest         float64
	TotalTime                 float64
	RequestsResults           []RequestResult
	StatusCodes               map[int]int
}
