package sgf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoSgf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoSgf Suite")
}
