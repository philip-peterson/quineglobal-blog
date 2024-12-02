// Package model has domain models used throughout the application.
package model

import (
	"time"

	g "maragu.dev/gomponents"
)

// Thing with a name.
type Thing struct {
	Name string
}

type QuinePost struct {
	Title   string
	Id      string // used as guid in rss, id in atom
	Updated time.Time
	Created time.Time
	Content []g.Node
}
