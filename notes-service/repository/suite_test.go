// +build unit integration

package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestSuiteNotesServiceRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSuiteNotesServiceRepository")
}
