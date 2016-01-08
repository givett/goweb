package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type teststruct struct {
	Test string
}

func handler(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t teststruct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Test)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8080
}
