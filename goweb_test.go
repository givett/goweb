package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "goweb"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Goweb", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
		server.AppendHandlers(Handler)
	})

	Describe("Testing handler with GET", func() {
		Context("Simulate a GET call", func() {
			It("Should return Hello World", func() {
				res, err := http.Get(server.URL())
				if err != nil {
					log.Fatal(err)
				}
				greeting, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", greeting)

				Ω(greeting).Should(Equal([]byte("Hello World!")))
			})
		})
	})

	Describe("Testing Handler with POST", func() {
		Context("Checks if true equals true", func() {
			It("True equal true?", func() {
				Expect(true).To(Equal(true))
			})
		})

		Context("Simulate a POST call", func() {
			It("Should return json string", func() {
				// Create book object
				testbook := Book{Isbn: "123-4567890123", Title: "Huckleberry Finn", Author: "Kernel Sanders", Price: 8.99}

				body, _ := json.Marshal(testbook)
				req, err := http.NewRequest("POST", server.URL()+"/book/", bytes.NewReader(body))

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("testbook.Title: %s\n", testbook.Title)

				Ω(testbook.Title).Should(Equal("Huckleberry Finn"))

				fmt.Printf("req.Body: %s\n", req.Body)
				fmt.Printf("res.Body: %s\n", res.Body)

				req.Body.Close()
				res.Body.Close()

			})
		})
	})
})
