//+build wireinject

package main

import (
	aws_config "github.com/geoffLondon/cdk-notes-api/aws"
	aws_dynamodb "github.com/geoffLondon/cdk-notes-api/aws/dynamodb"
	"github.com/geoffLondon/cdk-notes-api/configuration"
	service_repository "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	"github.com/geoffLondon/cdk-notes-api/resolver"
	"github.com/geoffLondon/cdk-notes-api/uuid"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	configuration.NewConfig,

	aws_config.NewAwsConfig,

	aws_dynamodb.NewDynamoDBFromConfig,
	aws_dynamodb.NewDynamoNotesServiceTable,
	aws_dynamodb.NewDefaultDynamoDbClient,
	wire.Bind(new(aws_dynamodb.DynamoDbClient), new(*aws_dynamodb.DefaultDynamoDbClient)),

	wire.Struct(new(DefaultContainer), "*"),
	wire.Bind(new(Container), new(DefaultContainer)),

	resolver.NewDefaultCreateNoteResolver,
	wire.Bind(new(resolver.CreateNoteResolver), new(*resolver.DefaultCreateNoteResolver)),

	service_repository.NewDynamoServiceRepository,
	wire.Bind(new(service_repository.ServiceRepository), new(*service_repository.DynamoServiceRepository)),

	resolver.NewValidatorImpl,
	wire.Bind(new(resolver.Validator), new(*resolver.ValidatorImpl)),

	uuid.NewDefaultUuidGenerator,
	wire.Bind(new(uuid.UuidGenerator), new(*uuid.DefaultUuidGenerator)),
)

func InitializeAppsyncContainer() (Container, error) {
	wire.Build(Set)
	return DefaultContainer{}, nil
}
