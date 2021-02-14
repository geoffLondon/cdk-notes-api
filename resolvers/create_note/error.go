package create_note_resolver

import "errors"

var (
	ErrMissingText = errors.New("note text missing")
	//ErrMissingDone = errors.New("done missing")
)
