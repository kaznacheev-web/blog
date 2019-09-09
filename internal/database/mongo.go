package database

import "github.com/kaznacheev-web/blog/internal/config"

// MongoDatabase is a MongoDB communication layer
type MongoDatabase struct {
}

// NewMongoDatabase creates a new database connection
func NewMongoDatabase(cfg config.MainConfig) (*MongoDatabase, error) {
	return &MongoDatabase{}, nil
}
