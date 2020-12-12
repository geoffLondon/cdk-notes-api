package service_resolver

import (
	"context"
	service_repository "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	"github.com/geoffLondon/cdk-notes-api/uuid"
	log "github.com/sirupsen/logrus"
)

type CreateNoteResolver interface {
	Handle(context.Context, CreateNoteParams) (string, error)
}

type CreateNoteParams struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type DefaultCreateNoteResolver struct {
	serviceRepository service_repository.ServiceRepository
	uuidGenerator     uuid.UuidGenerator
}

func NewDefaultCreateNoteResolver(serviceRepository service_repository.ServiceRepository, uuidGenerator uuid.UuidGenerator) *DefaultCreateNoteResolver {
	return &DefaultCreateNoteResolver{serviceRepository: serviceRepository, uuidGenerator: uuidGenerator}
}

func (resolver DefaultCreateNoteResolver) Handle(ctx context.Context, params CreateNoteParams) (string, error) {
	log.WithFields(log.Fields{"ctx": ctx, "params": params}).Info("request received")

	if err := params.validate(); err != nil {
		log.WithFields(log.Fields{"params": params, "err": err}).Warn("failed validating inputs")
		return "", err
	}

	service := service_repository.NotesService{
		Id:        resolver.uuidGenerator.New(),
		Name:      params.Name,
		Completed: params.Completed,
	}

	if err := resolver.serviceRepository.Save(ctx, service); err != nil {
		log.WithFields(log.Fields{"service": service, "err": err}).Warn("failed saving service")
		return "", err
	}

	return service.Id, nil
}

func (params CreateNoteParams) validate() error {
	if params.Id == "" {
		return ErrMissingId
	}
	if params.Name == "" {
		return ErrMissingName
	}
	if params.Completed == false {
		return ErrMissingCompleted
	}

	return nil
}
