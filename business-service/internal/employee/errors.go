package employee

import "errors"

// Repository level
var (
	ErrEmployeeNotFound = errors.New("employee not found")
)
