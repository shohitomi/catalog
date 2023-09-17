package repository

import "errors"

// ErrNotFound is returned when a requested item is not found
var ErrNotFound = errors.New("not found")
