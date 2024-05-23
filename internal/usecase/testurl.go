package usecase

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"renatonasc/stresstest/internal/entity"
	"sync"
	"time"
)

type TestUrlUseCase struct {
	RequestTest entity.RequestTest
}

func NewTestUrl(requestTest entity.RequestTest) *TestUrlUseCase {

	return &TestUrlUseCase{
		RequestTest: requestTest,
	}
}

func (u *TestUrlUseCase) Execute() (*entity.RequestTestResult, error) {

	// cria um canal para enviar os dados para os workers
	data := make(chan *entity.RequestResult)
	wg := sync.WaitGroup{}
	wg.Add(u.RequestTest.NumberOfRequests)
	startTime := time.Now()

	for i := 0; i < u.RequestTest.NumberOfConcurrentRequests; i++ {
		go worker(i, data, &wg)
	}
	statusCodes := make(map[int]int)
	requestsResults := make([]entity.RequestResult, u.RequestTest.NumberOfRequests)
	for i := 0; i < u.RequestTest.NumberOfRequests; i++ {
		requestsResults[i] = entity.RequestResult{
			Request: u.RequestTest.Request,
		}
		data <- &requestsResults[i]
	}
	wg.Wait()
	fmt.Println("All done")
	close(data)

	totalDuration := time.Duration(0)
	for _, r := range requestsResults {
		r.Duration = time.Duration(r.TimeEnd.Sub(r.TimeStart))
		totalDuration += r.Duration

		statusCodes[r.StatusCode]++

	}
	endTime := time.Now()

	// fmt.Printf("Total Duration: %f\n", totalDuration.Seconds())
	durationPerRequest := totalDuration.Seconds() / float64(u.RequestTest.NumberOfRequests)
	// fmt.Printf("Duration per request: %f\n", durationPerRequest)
	totalTime := endTime.Sub(startTime)

	return &entity.RequestTestResult{
		RequestTest:               u.RequestTest,
		NumberOfCompletedRequests: statusCodes[200],
		NumberOfFailedRequests:    u.RequestTest.NumberOfRequests - statusCodes[200],
		TimeMeanByRequest:         durationPerRequest,
		TotalTime:                 totalTime.Seconds(),
		RequestsResults:           requestsResults,
		StatusCodes:               statusCodes,
	}, nil
}

func worker(workerId int, data chan *entity.RequestResult, wg *sync.WaitGroup) {
	ctx := context.Background()
	client := &http.Client{
		Transport: &http.Transport{},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
	for x := range data {

		req, err := http.NewRequestWithContext(ctx, x.Request.Method, x.Request.Url, nil)
		if err != nil {
			x.Body = "Erro on create request"
			wg.Done()
			continue
		}
		x.TimeStart = time.Now()
		resp, err := client.Do(req)
		if err != nil {
			x.TimeEnd = time.Now()
			x.StatusCode = 0
			x.Body = fmt.Sprintf("Error on make request: %v", err)
			fmt.Println(x.Body)
			wg.Done()
			continue
		}
		defer resp.Body.Close()
		x.TimeEnd = time.Now()
		x.StatusCode = resp.StatusCode

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			x.Body = "Erro on read body"
			wg.Done()
			continue
		}
		x.Body = string(body)
		wg.Done()
		continue
	}
}
