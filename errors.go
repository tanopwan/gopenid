package gopenid

import (
	"errors"
)

var (
	ErrJWTHeaderMissingKID = errors.New("expecting JWT header to have string kid")
	ErrPublicKeyIsNotFound = errors.New("unable to find public key (key is too old)")
)
