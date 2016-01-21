package main_test

import (
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

	Describe("Testing handler", func() {

		It("should be get returns hello world", func() {
			res, err := http.Get(server.URL())
			if err != nil {
				log.Fatal(err)
			}
			greeting, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s", greeting)

			Ω(greeting).Should(Equal([]byte("Hello World!")))
		})
	})
})
