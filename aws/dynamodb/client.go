package aws_dynamodb

import (
	"context"
	"github.com/guregu/dynamo"
)

type DynamoDbClient interface {
	TableName() string
	Put(ctx context.Context, data interface{}) error
	FindAll(ctx context.Context, out interface{}) error
}

type DefaultDynamoDbClient struct {
	table dynamo.Table
}

func NewDefaultDynamoDbClient(table dynamo.Table) *DefaultDynamoDbClient {
	return &DefaultDynamoDbClient{
		table: table,
	}
}

func (client *DefaultDynamoDbClient) TableName() string {
	return client.table.Name()
}

func (client *DefaultDynamoDbClient) Put(ctx context.Context, data interface{}) error {
	return client.table.Put(data).RunWithContext(ctx)
}

func (client *DefaultDynamoDbClient) FindAll(ctx context.Context, out interface{}) error {
	return client.table.Scan().AllWithContext(ctx, out)
}
