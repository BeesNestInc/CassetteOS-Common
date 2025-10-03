package version

import (
	"errors"
	_ "github.com/mattn/go-sqlite3" // nolint
)

var (
	ErrLegacyVersionNotFound = errors.New("legacy version not found")
	ErrVersionNotFound       = errors.New("version (non-legacy) not found")
)