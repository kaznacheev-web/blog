package models

import (
	"time"

	"golang.org/x/text/language"
)

// Language is an language code
type Language string

const (
	// EN English
	EN Language = "EN"
	// DE German
	DE Language = "DE"
	// RU Russian
	RU Language = "RU"
)

// Article describes an article entity
type Article struct {
	Title     string
	Preview   string
	Text      string
	Slug      string
	Published bool
	PubDate   time.Time
	Language  language.Language
	Tags      []string
}

// Talk describes a talk entity
type Talk struct {
	Title     string
	Text      string
	VideoURL  *string
	SlideURL  *string
	Slug      string
	Published bool
	PubDate   time.Time
	Language  Language
	Tags      []string
}

// SimplePage describes a page with a free text structure
type SimplePage struct {
	Name  string
	Title string
	Text  string
	Slug  string
}
