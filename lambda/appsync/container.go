package main

import (
	app_sync_resolvers "github.com/geoffLondon/aws-appsync-resolvers"
	"github.com/geoffLondon/cdk-notes-api/resolvers/create_note"
	"github.com/geoffLondon/cdk-notes-api/resolvers/list_notes"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Resolver() app_sync_resolvers.Repository
}

type DefaultContainer struct {
	CreateNoteResolver create_note_resolver.CreateNoteResolver
	ListNotesResolver  list_notes_resolver.ListNotesResolver
}

func (container DefaultContainer) Resolver() app_sync_resolvers.Repository {
	repository := app_sync_resolvers.New()

	resolversMap := map[string]interface{}{
		"mutation.createNote": container.CreateNoteResolver.Handle,
		"query.listNotes":     container.ListNotesResolver.Handle,
	}

	for resolver, handler := range resolversMap {
		if err := repository.Add(resolver, handler); err != nil {
			log.WithField("err", err).Warn("error adding resolvers to repository")
		}
		log.WithField("resolvers", resolver).Info("resolvers added to repository")
	}

	return repository
}
