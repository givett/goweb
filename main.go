package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))

		rBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading app request, %v", rBody)
		}

		err = json.Unmarshal(rBody, r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error unmarshaling app %v", err)
		}

		// create the statement string
		//var sStmt = "insert into test (gopher_id, created) values ($1, $2)"

		// lazily open db (doesn't truly open until first request)
		//db, err := sql.Open("postgres", "host=localhost dbname=testdb sslmode=disable")
		//if err != nil {
		//log.Fatal(err)
		//}

		//stmt, err := db.Prepare(sStmt)
		//if err != nil {
		//log.Fatal(err)
		//}

		//fmt.Printf("StartTime: %v\n", time.Now())

		//res, err := stmt.Exec(1, time.Now())
		//if err != nil || res == nil {
		//log.Fatal(err)
		//}

		// close statement
		//stmt.Close()

		// close db
		//db.Close()

		//fmt.Printf("StopTime: %v\n", time.Now())

		fmt.Fprintf(w, "rBody is, %s", rBody)

	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
	// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8080
}
