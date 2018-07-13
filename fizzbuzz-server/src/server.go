package main

import (
	"fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
    port := ":8020"

    fmt.Printf("Starting FizzBuzz server on %s \n", port)

    router := mux.NewRouter()

    router.HandleFunc("/fizzbuzz/{max}", PrintFizzBuzz).Methods("GET")

    log.Fatal(http.ListenAndServe(port, router))
}