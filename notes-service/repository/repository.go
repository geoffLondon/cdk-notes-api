package repository

import (
	"context"
	"github.com/geoffLondon/cdk-notes-api/aws/dynamodb"
	"github.com/guregu/dynamo"
	log "github.com/sirupsen/logrus"
)

type ServiceRepository interface {
	FindAll(ctx context.Context) ([]NotesService, error)
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

func (repo *DynamoServiceRepository) FindAll(ctx context.Context) ([]NotesService, error) {
	log.Debug("Finding all Notes")
	notes := make([]NotesService, 0)
	if err := repo.dynamoDbClient.FindAll(ctx, &notes); err != nil {
		if err == dynamo.ErrNotFound {
			log.Debug("No items found in table")
		} else {
			return []NotesService{}, err
		}
	}
	return notes, nil
}
