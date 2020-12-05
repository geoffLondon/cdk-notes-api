package aws_dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/guregu/dynamo"
)

func NewDynamoDBFromConfig(config *aws.Config) (*dynamo.DB, error) {
	dynamoSessionOptions := newDynamoDbSessionOptions(config)
	dynamoSession, err := session.NewSessionWithOptions(dynamoSessionOptions)
	if err != nil {
		return nil, err
	}
	return newDynamoClient(dynamoSession), nil
}

func newDynamoClient(session *session.Session) *dynamo.DB {
	db := dynamodb.New(session)
	xray.AWS(db.Client)
	return dynamo.NewFromIface(db)
}

func newDynamoDbSessionOptions(config *aws.Config) session.Options {
	return session.Options{
		Config: *config,
	}
}
