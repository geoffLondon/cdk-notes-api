package resolver

import (
	"context"
	service_repository "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	"github.com/geoffLondon/cdk-notes-api/uuid"
	log "github.com/sirupsen/logrus"
)

type CreateNoteResolver interface {
	Handle(ctx context.Context, noteRequest NoteRequest) (bool, error)
}

type DefaultCreateNoteResolver struct {
	serviceRepository service_repository.ServiceRepository
	validator         Validator
	uuidGenerator     uuid.UuidGenerator
}

func NewDefaultCreateNoteResolver(serviceRepository service_repository.ServiceRepository, validator Validator, uuidGenerator uuid.UuidGenerator) *DefaultCreateNoteResolver {
	return &DefaultCreateNoteResolver{
		serviceRepository: serviceRepository,
		validator:         validator,
		uuidGenerator:     uuidGenerator,
	}
}

func (resolver DefaultCreateNoteResolver) Handle(ctx context.Context, noteRequest NoteRequest) (bool, error) {
	log.WithFields(log.Fields{"noteRequest": noteRequest}).Info("note request received")

	if err := resolver.validator.Validate(noteRequest); err != nil {
		return false, err
	}

	service := service_repository.NotesService{
		Id:        resolver.uuidGenerator.New(),
		Name:      noteRequest.Name,
		Completed: noteRequest.Completed,
	}

	if err := resolver.serviceRepository.Save(ctx, service); err != nil {
		log.WithFields(log.Fields{"service": service, "err": err}).Warn("failed saving service to db")
		return false, err
	}

	return true, nil
}
