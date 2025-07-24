package models

import (
	"errors"
)

// This is similar to pgx.ErrNoRows, but this is for application layer/client
// This provides an abstraction. pgx.ErrNoRows might have too uch info for our frontend devs
// We simplify it like so
var ErrNoRecord = errors.New("models: no matching record found")
