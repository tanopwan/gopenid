package gopenid

import (
	"github.com/pkg/errors"
)

var (
	// ErrExternalDAO calling external dao error without business reason
	ErrExternalDAO = errors.New("external dao error")

	// ErrInternalServerError json error, parsing error
	ErrInternalServerError = errors.New("internal server error")
)
