package main

import (
	"github.com/geoffLondon/aws-appsync-resolvers"
	"github.com/geoffLondon/cdk-notes-api/resolver"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Resolver() resolvers.Repository
}

type DefaultContainer struct {
	CreateNoteResolver resolver.CreateNoteResolver
}

func (container DefaultContainer) Resolver() resolvers.Repository {
	repository := resolvers.New()

	resolversMap := map[string]interface{}{
		"mutation.createNote": container.CreateNoteResolver.Handle,
	}

	for resolver, handler := range resolversMap {
		if err := repository.Add(resolver, handler); err != nil {
			log.WithField("err", err).Warn("error adding resolver to repository")
		}
		log.WithField("resolver", resolver).Info("resolver added to repository")
	}

	return repository
}
