package library_test

import (
	//. "main"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Library", func() {
	Context("Should Have Books", func() {
		It("Should have mystery books", func() {
			//TEST CODE
			// Do an HTTP request. x represents the return CODE
			// Expect(x).To(Equal(200))
			Expect("itisamystery").To(Equal("itisamystery"))
			Ω("itisamystery").Should(Equal("itisamystery"))
		})
		It("Should have non-fiction books", func() {
			//TEST CODE
		})
	})

	It("Should have librarians", func() {
		//TEST CODE
	})
})
