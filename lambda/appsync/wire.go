//+build wireinject

package main

import (
	"github.com/geoffLondon/cdk-notes-api/aws"
	"github.com/geoffLondon/cdk-notes-api/aws/dynamodb"
	"github.com/geoffLondon/cdk-notes-api/configuration"
	"github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	"github.com/geoffLondon/cdk-notes-api/resolvers/create_note"
	"github.com/geoffLondon/cdk-notes-api/resolvers/list_notes"
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

	create_note_resolver.NewDefaultCreateNoteResolver,
	wire.Bind(new(create_note_resolver.CreateNoteResolver), new(*create_note_resolver.DefaultCreateNoteResolver)),

	list_notes_resolver.NewDefaultListNotesResolver,
	wire.Bind(new(list_notes_resolver.ListNotesResolver), new(*list_notes_resolver.DefaultListNotesResolver)),

	repository.NewDynamoServiceRepository,
	wire.Bind(new(repository.ServiceRepository), new(*repository.DynamoServiceRepository)),

	create_note_resolver.NewValidatorImpl,
	wire.Bind(new(create_note_resolver.Validator), new(*create_note_resolver.ValidatorImpl)),

	uuid.NewDefaultUuidGenerator,
	wire.Bind(new(uuid.UuidGenerator), new(*uuid.DefaultUuidGenerator)),
)

func InitializeAppsyncContainer() (Container, error) {
	wire.Build(Set)
	return DefaultContainer{}, nil
}
