package aws_dynamodb

import (
	"github.com/geoffLondon/cdk-notes-api/configuration"
	"github.com/guregu/dynamo"
)

func NewDynamoNotesServiceTable(dynamoDB *dynamo.DB, configuration configuration.Configuration) dynamo.Table {
	return dynamoDB.Table(configuration.GeoffCdkNotesTableName)
}
