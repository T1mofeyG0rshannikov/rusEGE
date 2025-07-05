package exceptions

import (
	"errors"
)

var ErrInvalidJwtToken = errors.New("invalid JWT token")
var ErrNoAuthHeader = errors.New("no Auhorization header")