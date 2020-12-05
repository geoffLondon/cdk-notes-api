//+build wireinject

package main

import (
	"github.com/google/wire"
	aws_config "cdk-notes-api/aws"
	aws_dynamodb "cdk-notes-api/aws/dynamodb"
	"cdk-notes-api/configuration"
	service_repository "cdk-notes-api/notes-service/repository"
	"cdk-notes-api/resolver/service"
	"cdk-notes-api/uuid"
)

var Set = wire.NewSet(
	// Configuration
	configuration.NewConfig,

	// AWS clients & resources
	aws_config.NewAwsConfig,

	aws_dynamodb.NewDynamoDBFromConfig,
	aws_dynamodb.NewDynamoCallFlowServiceTable,
	aws_dynamodb.NewDefaultDynamoDbClient,
	wire.Bind(new(aws_dynamodb.DynamoDbClient), new(*aws_dynamodb.DefaultDynamoDbClient)),

	// Container and resolvers
	wire.Struct(new(DefaultContainer), "*"),
	wire.Bind(new(Container), new(DefaultContainer)),

	service_resolver.NewDefaultCreateNoteResolver,
	wire.Bind(new(service_resolver.CreateNoteResolver), new(*service_resolver.DefaultCreateNoteResolver)),

	// Repos
	service_repository.NewDynamoServiceRepository,
	wire.Bind(new(service_repository.ServiceRepository), new(*service_repository.DynamoServiceRepository)),

	//	Services
	uuid.NewDefaultUuidGenerator,
	wire.Bind(new(uuid.UuidGenerator), new(*uuid.DefaultUuidGenerator)),
)

func InitializeAppsyncContainer() (Container, error) {
	wire.Build(Set)
	return DefaultContainer{}, nil
}
