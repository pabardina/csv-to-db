package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTest Suite")
}
