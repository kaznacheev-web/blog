package database

import (
	"github.com/kaznacheev-web/blog/internal/config"
	"github.com/kaznacheev-web/blog/internal/models"
)

// MongoDatabase is a MongoDB communication layer
type MongoDatabase struct {
}

// NewMongoDatabase creates a new database connection
func NewMongoDatabase(cfg config.Database) (*MongoDatabase, error) {
	return &MongoDatabase{}, nil
}

func (mdb *MongoDatabase) GetArticles(page int) ([]models.Article, error) {

}

func (mdb *MongoDatabase) GetArticle(slug string) (*models.Article, error) {

}

func (mdb *MongoDatabase) GetArticleCount() (int, error) {

}

func (mdb *MongoDatabase) GetTalks(page int) ([]models.Talk, error) {

}

func (mdb *MongoDatabase) GetTalk(slug string) (*models.Talk, error) {

}

func (mdb *MongoDatabase) GetTalkCount() (int, error) {

}

func (mdb *MongoDatabase) GetSimplePage(key string) (*models.SimplePage, error) {

}
