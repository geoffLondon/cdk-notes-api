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
	Handle(ctx context.Context, noteRequest NoteRequest) (NoteResponse, error)
}

type DefaultCreateNoteResolver struct {
	serviceRepository service_repository.ServiceRepository
	validator         Validator
	//uuidGenerator     uuid.UuidGenerator
}

func NewDefaultCreateNoteResolver(serviceRepository service_repository.ServiceRepository, validator Validator) *DefaultCreateNoteResolver {
	return &DefaultCreateNoteResolver{
		serviceRepository: serviceRepository,
		validator:         validator,
	}
}

func (resolver DefaultCreateNoteResolver) Handle(ctx context.Context, noteRequest NoteRequest) (NoteResponse, error) {
	log.WithFields(log.Fields{"noteRequest": noteRequest}).Info("note request received")

	/*	if noteRequest.Id == "" {
		log.WithFields(log.Fields{"noteId": noteRequest.Id}).Warn("note id missing, still!")
		return "", errors.New("error, missing fields")
	}*/

	if err := resolver.validator.Validate(noteRequest); err != nil {
		return NoteResponse{}, err
	}

	service := service_repository.NotesService{
		//Id:        resolver.uuidGenerator.New(),
		Id:        noteRequest.Id,
		Name:      noteRequest.Name,
		Completed: noteRequest.Completed,
	}

	if err := resolver.serviceRepository.Save(ctx, service); err != nil {
		log.WithFields(log.Fields{"service": service, "err": err}).Warn("failed saving service to db")
		return NoteResponse{}, err
	}

	response := NoteResponse{
		Note: Note{
			Id:        service.Id,
			Name:      service.Name,
			Completed: service.Completed,
		},
	}

	return response, nil
}
