package common

import "errors"

var (
	ErrDatabaseConnection  = errors.New("failed to connect to db")
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckViolation      = errors.New("check violation")
)
