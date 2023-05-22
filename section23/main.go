package main

import (
	"fmt"
	"log"
	"os"
)

type Result struct {
	Response *os.File
	Error    error
}

func CheckFiles(done <-chan interface{}, filenames ...string) <-chan Result {
	results := make(chan Result)

	go func() {
		defer close(results)

		for _, filename := range filenames {
			var result Result

			file, err := os.Open(filename)
			result = Result{file, err}

			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}

func main() {
	done := make(chan interface{})

	defer close(done)

	filenames := []string{"main.go", "x.go"}

	for result := range CheckFiles(done, filenames...) {
		if result.Error != nil {
			log.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Name())
	}
}
