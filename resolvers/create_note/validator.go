package create_note_resolver

import "errors"

type Validator interface {
	Validate(noteRequest NoteRequest) error
}

type ValidatorImpl struct {
}

func NewValidatorImpl() *ValidatorImpl {
	return &ValidatorImpl{}
}

const (
	ErrMissingText = "note text missing"
)

func (validator ValidatorImpl) Validate(noteRequest NoteRequest) error {
	if noteRequest.Text == "" {
		return errors.New(ErrMissingText)
	}

	return nil
}
