// +build unit

package create_note_resolver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCallStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateNoteResolverSuite")
}
