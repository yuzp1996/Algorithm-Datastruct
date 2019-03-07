package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinklist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linklist Suite")
}
