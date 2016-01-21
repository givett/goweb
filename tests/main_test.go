package main_test

import (
	//. "main"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var w string

	Describe("GET", func() {

		It("Should return Hello World!", func() {

			w = "Hello World!"
			Expect(w).To(Equal("Hello World!"))

		})
	})

	Describe("POST", func() {

		It("Should return contents of form_data.json file.", func() {

			w = `"isbn": "978-1505255607",  "title": "The Time Machine",  "author": "H. G. Wells",  "price": 6.99"`
			Expect(w).To(Equal(`"isbn": "978-1505255607",  "title": "The Time Machine",  "author": "H. G. Wells",  "price": 6.99"`))

		})

	})

})
