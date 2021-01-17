package repository

import (
	"context"
	aws_dynamodb "github.com/geoffLondon/cdk-notes-api/aws/dynamodb"
	log "github.com/sirupsen/logrus"
)

type ServiceRepository interface {
	Save(ctx context.Context, service NotesService) error
}

type DynamoServiceRepository struct {
	dynamoDbClient aws_dynamodb.DynamoDbClient
}

func NewDynamoServiceRepository(dynamoDbClient aws_dynamodb.DynamoDbClient) *DynamoServiceRepository {
	return &DynamoServiceRepository{
		dynamoDbClient: dynamoDbClient,
	}
}

func (repo *DynamoServiceRepository) Save(ctx context.Context, notesService NotesService) error {
	log.WithFields(log.Fields{"service": notesService}).Info("Saving request")

	err := repo.dynamoDbClient.Put(ctx, &notesService)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error writing to Dynamo")
	} else {
		log.WithFields(log.Fields{"table": repo.dynamoDbClient.TableName()}).Info("Saved record successfully")
	}

	return err
}

// TODO need to add Get() function
