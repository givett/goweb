package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoweb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goweb Suite")
}
