package aws_dynamodb

import (
	"cdk-notes-api/configuration"
	"github.com/guregu/dynamo"
)

func NewDynamoCallFlowServiceTable(dynamoDB *dynamo.DB, configuration configuration.Configuration) dynamo.Table {
	return dynamoDB.Table(configuration.GeoffCdkNotesTableName)
}
