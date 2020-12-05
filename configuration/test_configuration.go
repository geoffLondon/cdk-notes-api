// +build unit integration

package configuration

import (
	"github.com/resilientplc/protect-call-flow-service/fixtures"
)

func NewTestConfiguration() Configuration {
	return Configuration{
		CallFlowServicesTableName: fixtures.CallFlowServicesTableName,
	}
}
