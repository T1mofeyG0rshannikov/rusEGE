package exceptions

import (
	"errors"
)

var ErrIncorrectPassword = errors.New("incorrect password")
var ErrUserNotFound = errors.New("user not found")
var ErrUsernameExist = errors.New("username already exists")
