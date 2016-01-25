package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

// Book comment
type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Handler comment
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		//fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))

	} else if r.Method == "POST" {
		fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))

		rBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading app request, %v", rBody)
		}

		var book Book

		err = json.Unmarshal(rBody, &book)
		if err != nil {
			fmt.Fprintf(w, "Error unmarshaling app %v", err)
		}

		fmt.Fprintf(w, "title is %s", book.Title)

	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Print("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
	// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8080
	// curl -X POST -d @form_data.json http://localhost:8080 -v
	// curl localhost:8080 -X GET
	// curl :8080 -X POST -d @form_data.json
}
