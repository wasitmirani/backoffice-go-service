package database

import "errors"

var (
	ErrUnsupportedDriver = errors.New("unsupported database driver")
	ErrConnectionFailed  = errors.New("database connection failed")
	ErrDriverNotFound    = errors.New("database driver not found")
	ErrNotConnected      = errors.New("database not connected")
)

