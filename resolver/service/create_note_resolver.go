package service_resolver

import (
	"context"
	service_repository "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	//"github.com/geoffLondon/cdk-notes-api/uuid"
	log "github.com/sirupsen/logrus"
)

type CreateNoteParams struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type CreateNoteResolver interface {
	Handle(context.Context, CreateNoteParams) (string, error)
}

type DefaultCreateNoteResolver struct {
	serviceRepository service_repository.ServiceRepository
	//uuidGenerator     uuid.UuidGenerator
}

func NewDefaultCreateNoteResolver(serviceRepository service_repository.ServiceRepository) *DefaultCreateNoteResolver {
	return &DefaultCreateNoteResolver{serviceRepository: serviceRepository}
}

func (resolver DefaultCreateNoteResolver) Handle(ctx context.Context, createNoteParams CreateNoteParams) (string, error) {
	log.WithFields(log.Fields{"ctx": ctx, "createNoteParams": createNoteParams}).Info("request received")

	if err := createNoteParams.validate(); err != nil {
		log.WithFields(log.Fields{"createNoteParams": createNoteParams, "err": err}).Warn("failed validating inputs")
		return "", err
	}

	service := service_repository.NotesService{
		//Id:        resolver.uuidGenerator.New(),
		Id:        createNoteParams.Id,
		Name:      createNoteParams.Name,
		Completed: createNoteParams.Completed,
	}

	log.WithFields(log.Fields{"=======serviceId========": service.Id, "=======serviceName========": service.Name, "=======serviceCompleted========": service.Completed, "=======service========": service}).Info("NotesService")

	if err := resolver.serviceRepository.Save(ctx, service); err != nil {
		log.WithFields(log.Fields{"service": service, "err": err}).Warn("failed saving service")
		return "", err
	}

	return service.Id, nil
}

func (createNoteParams CreateNoteParams) validate() error {
	if createNoteParams.Id == "" {
		return ErrMissingId
	}
	if createNoteParams.Name == "" {
		return ErrMissingName
	}
	if createNoteParams.Completed == false {
		return ErrMissingCompleted
	}

	return nil
}
