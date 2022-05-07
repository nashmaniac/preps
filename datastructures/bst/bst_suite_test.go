package bst_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBst(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bst Suite")
}
