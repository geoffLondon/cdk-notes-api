package main

import (
	"github.com/geoffLondon/appsync-resolvers"
	"cdk-notes-api/resolver/service"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Resolver() resolvers.Repository
}

type DefaultContainer struct {
	CreateServiceResolver   service_resolver.CreateNoteResolver
}

func (container DefaultContainer) Resolver() resolvers.Repository {
	repository := resolvers.New()

	resolversMap := map[string]interface{}{
		"mutation.createService":             container.CreateServiceResolver.Handle,
	}

	for resolver, handler := range resolversMap {
		if err := repository.Add(resolver, handler); err != nil {
			log.WithField("err", err).Warn("error add resolver to repository")
		}
	}

	return repository
}
