package food_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFood(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Food Suite")
}
