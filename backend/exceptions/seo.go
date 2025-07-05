package exceptions

import (
	"errors"
)

var ErrIndexSeoAlreadyExists = errors.New("index seo object already exists")
var ErrIndexSeoNotFound = errors.New("index seo not found")