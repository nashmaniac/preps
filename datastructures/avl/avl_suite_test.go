package avl_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAvl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Avl Suite")
}
