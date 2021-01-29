package main

import (
	app_sync_resolvers "github.com/geoffLondon/aws-appsync-resolvers"
	create_note_resolver "github.com/geoffLondon/cdk-notes-api/resolvers/create_note"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Resolver() app_sync_resolvers.Repository
}

type DefaultContainer struct {
	CreateNoteResolver create_note_resolver.CreateNoteResolver
}

func (container DefaultContainer) Resolver() app_sync_resolvers.Repository {
	repository := app_sync_resolvers.New()

	resolversMap := map[string]interface{}{
		"mutation.createNote": container.CreateNoteResolver.Handle,
	}

	for res, handler := range resolversMap {
		if err := repository.Add(res, handler); err != nil {
			log.WithField("err", err).Warn("error adding resolvers to repository")
		}
		log.WithField("resolvers", res).Info("resolvers added to repository")
	}

	return repository
}
