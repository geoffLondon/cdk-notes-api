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
	return &DynamoServiceRepository{dynamoDbClient: dynamoDbClient}
}

func (repo *DynamoServiceRepository) Save(ctx context.Context, service NotesService) error {
	log.WithFields(log.Fields{"ctx": ctx, "service": service}).Info("Save request")

	return repo.dynamoDbClient.Put(ctx, &service)
}
