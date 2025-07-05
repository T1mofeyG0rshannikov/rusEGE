package exceptions


import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task with with number already exists")