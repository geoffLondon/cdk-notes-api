package create_note_resolver

type Validator interface {
	Validate(noteRequest NoteRequest) error
}

type ValidatorImpl struct {
}

func NewValidatorImpl() *ValidatorImpl {
	return &ValidatorImpl{}
}

func (validator ValidatorImpl) Validate(noteRequest NoteRequest) error {
	if noteRequest.Text == "" {
		return ErrMissingText
	}
	/*	if validator.Done == false {
		return ErrMissingDone
	}*/

	return nil
}
