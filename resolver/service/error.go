package service_resolver

import "errors"

var (
	ErrMissingId  = errors.New("customer id missing")
	ErrMissingName = errors.New("service name missing")
	ErrMissingCompleted   = errors.New("completed missing")
)
