package list_notes_resolver

import (
	"context"
	"github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	log "github.com/sirupsen/logrus"
)

type ListNotesResolver interface {
	Handle(ctx context.Context) ([]repository.NotesService, error)
}

type DefaultListNotesResolver struct {
	serviceRepository repository.ServiceRepository
}

func NewDefaultListNotesResolver(serviceRepository repository.ServiceRepository) *DefaultListNotesResolver {
	return &DefaultListNotesResolver{
		serviceRepository: serviceRepository,
	}
}

func (rcvr *DefaultListNotesResolver) Handle(ctx context.Context) ([]repository.NotesService, error) {
	notes, err := rcvr.serviceRepository.FindAll(ctx)
	if err != nil {
		log.WithField("err", err).Warn("failed while scanning db for notes")
		return []repository.NotesService{}, err
	}
	return notes, nil
}
