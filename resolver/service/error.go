package service_resolver

import "errors"

var (
	ErrMissingId        = errors.New("note id missing, oops")
	ErrMissingName      = errors.New("service name missing")
	ErrMissingCompleted = errors.New("completed missing")
)
