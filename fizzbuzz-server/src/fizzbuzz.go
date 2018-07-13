package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
)

// PrintFizzBuzz : Sends request to GoRoutine to do the fizz and buzzes
func PrintFizzBuzz(writer http.ResponseWriter, reader *http.Request) {
	params := mux.Vars(reader)

	max, err := strconv.ParseInt(params["max"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	}

	writer.Write([]byte(FizzBuzz(max)))
}

// FizzBuzz : Computes Fizzes and Buzzes
func FizzBuzz(max int64) string {
	results := ""
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for i := int64(1); i <= max; i++ {
			iter := ""
			if i%3 == 0 {
				iter += "Fizz"
			}
			if i%5 == 0 {
				iter += "Buzz"
			}
			if iter == "" {
				results += fmt.Sprintf("%d", i) + "\n"
			} else {
				results += iter + "\n"
			}
		}
		defer wg.Done()
	}()

	wg.Wait()
	return results
}
