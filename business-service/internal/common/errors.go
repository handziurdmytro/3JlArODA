package common

import "errors"

var (
	ErrDatabaseConnection = errors.New("failed to connect to db")
)
