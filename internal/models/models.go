package models

import (
	"time"

	// "golang.org/x/text/language"
	"github.com/globalsign/mgo/bson"
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

type MongoMeta struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"-"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type TextEntity struct {
	Title    string   `json:"title"`
	Slug     string   `json:"slug"`
	Text     string   `json:"text"`
	Language Language `json:"language"`
}

// Article describes an article entity
type Article struct {
	MongoMeta  `bson:",inline"`
	TextEntity `bson:",inline"`
	Preview    string    `json:"preview"`
	Published  bool      `json:"published"`
	PubDate    time.Time `json:"publication_date"`
	Tags       []string  `json:"tags,omitempty"`
}

// Talk describes a talk entity
type Talk struct {
	MongoMeta
	TextEntity
	VideoURL  *string   `json:"video_url,omitempty"`
	SlideURL  *string   `json:"slide_url,omitempty"`
	Published bool      `json:"published"`
	PubDate   time.Time `json:"publication_date"`
	// Tags is a list of talk tags
	Tags []string `json:"tags,omitempty"`
}

// SimplePage describes a page with a free text structure
type SimplePage struct {
	MongoMeta
	TextEntity
	// PageKey is a unique key to find this single page
	PageKey string `json:"page_key"`
}
