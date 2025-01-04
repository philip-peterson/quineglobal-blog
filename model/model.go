// Package model has domain models used throughout the application.
package model

import (
	"time"

	g "maragu.dev/gomponents"
)

type QuinePost struct {
	Title       string
	Id          string // used as guid in rss, id in atom
	Updated     time.Time
	Created     time.Time
	Teaser      string
	Content     []g.Node
	FooterSegue string
}
