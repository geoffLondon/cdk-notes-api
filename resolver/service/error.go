package service_resolver

import "errors"

var (
	ErrMissingId        = errors.New("note id missing, ffs")
	ErrMissingName      = errors.New("note name missing")
	ErrMissingCompleted = errors.New("completed missing")
)
