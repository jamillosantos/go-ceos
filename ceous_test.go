package ceous_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/jamillosantos/macchiato"
)

func TestCeous(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "Ceous Test Suite")
}
