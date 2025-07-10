package exceptions


import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task with with number already exists")
var ErrRuleNotFound = errors.New("rule not found")
var ErrRecordNotFound = errors.New("record not found")
var ErrWordAlreadyExists = errors.New("word already exists")