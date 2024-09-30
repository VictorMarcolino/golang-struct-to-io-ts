package usecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsecaseSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "...")
}
