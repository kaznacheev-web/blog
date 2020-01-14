package database

import "errors"

var (
	ErrNotFound = errors.New("database not found")
	ErrGeneric  = errors.New("database error")
)
