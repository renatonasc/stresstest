package cmd

import (
	"fmt"
	"os"
	"renatonasc/stresstest/internal/entity"
	"renatonasc/stresstest/internal/usecase"
	"strings"

	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "StressTest",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: runRoot(),
}

func runRoot() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		// name, _ := cmd.Flags().GetString("name")
		// description, _ := cmd.Flags().GetString("description")
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return err
		}
		if url == "" {
			return fmt.Errorf("url is required")
		}
		numberOfRequests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		method, _ := cmd.Flags().GetString("method")
		method = strings.ToUpper(method)
		if method == "" {
			method = "GET"
		}

		// fmt.Printf("Name: %s\n", name)
		// fmt.Printf("Description: %s\n", description)
		fmt.Printf("Number of requests: %d\n", numberOfRequests)
		fmt.Printf("Concurrency: %d\n", concurrency)
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Method: %s\n", method)

		request := entity.Request{
			Url:    url,
			Method: method,
		}

		requestTest := entity.RequestTest{
			Request:                    request,
			NumberOfRequests:           numberOfRequests,
			NumberOfConcurrentRequests: concurrency,
			NumberOfRequestsPerSecond:  1,
		}

		testUrlUseCase := usecase.NewTestUrl(requestTest)
		result, err := testUrlUseCase.Execute()
		if err != nil {
			return err
		}

		// fmt.Printf("Result: %v\n", result)
		fmt.Printf("Number of completed requests: %d\n", result.NumberOfCompletedRequests)
		fmt.Printf("Number of failed requests: %d\n", result.NumberOfFailedRequests)
		fmt.Printf("Time mean by request: %f\n", result.TimeMeanByRequest)
		totalSeconds := int(result.TotalTime)
		minutes := totalSeconds / 60
		seconds := totalSeconds % 60
		fmt.Printf("Total time: %d minutes %d seconds\n", minutes, seconds)
		fmt.Printf("Status codes:\n")
		for k, v := range result.StatusCodes {
			fmt.Printf("Status code %d: %d requests\n", k, v)
		}
		return nil
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP("requests", "n", 1, "Number of requests to perform")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Number of multiple requests to make at a time")
	rootCmd.Flags().StringP("method", "m", "", "Method to use for the requests")
	rootCmd.Flags().StringP("url", "u", "", "URL to make the requests")
}
