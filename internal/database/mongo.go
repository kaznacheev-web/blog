package database

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kaznacheev-web/blog/internal/config"
	"github.com/kaznacheev-web/blog/internal/models"
)

const (
	pageSize = 10
)

const (
	collectionArticles = "articles"
	collectionTalks    = "talks"
	collectionPages    = "pages"
)

// MongoDatabase is a MongoDB communication layer
type MongoDatabase struct {
	*mongo.Database
}

// NewMongoDatabase creates a new database connection
func NewMongoDatabase(cfg config.Database) (*MongoDatabase, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("personal")
	return &MongoDatabase{
		Database: db,
	}, nil
}

func (mdb *MongoDatabase) GetArticles(ctx context.Context, page int) (res []models.Article, totalPages int, err error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64(pageSize * page))

	cur, err := mdb.Collection(collectionArticles).Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, 0, wrapErr(err)
	}

	err = cur.All(ctx, &res)
	if err != nil {
		return nil, 0, wrapErr(err)
	}

	total, err := mdb.GetArticleCount(ctx)
	totalPages = int(math.Ceil(float64(total) / pageSize))
	err = wrapErr(err)
	return
}

func (mdb *MongoDatabase) GetArticle(ctx context.Context, slug string) (res *models.Article, err error) {
	filter := bson.D{
		{"slug", slug},
	}
	res = new(models.Article)

	err = mdb.Collection(collectionArticles).FindOne(ctx, filter).Decode(res)
	err = wrapErr(err)
	return
}

func (mdb *MongoDatabase) GetArticleCount(ctx context.Context) (int, error) {
	c, err := mdb.Collection(collectionArticles).CountDocuments(ctx, bson.D{})
	return int(c), wrapErr(err)
}

func (mdb *MongoDatabase) GetTalks(ctx context.Context, page int) (res []models.Talk, totalPages int, err error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize * page))
	findOptions.SetSkip(int64(pageSize))

	cur, err := mdb.Collection(collectionTalks).Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, 0, wrapErr(err)
	}

	err = cur.All(ctx, &res)
	if err != nil {
		return nil, 0, wrapErr(err)
	}

	total, err := mdb.GetTalkCount(ctx)
	totalPages = int(math.Ceil(float64(total) / pageSize))
	err = wrapErr(err)
	return
}

func (mdb *MongoDatabase) GetTalk(ctx context.Context, slug string) (res *models.Talk, err error) {
	filter := bson.D{
		{"slug", slug},
	}
	res = new(models.Talk)

	err = mdb.Collection(collectionTalks).FindOne(ctx, filter).Decode(res)
	err = wrapErr(err)
	return
}

func (mdb *MongoDatabase) GetTalkCount(ctx context.Context) (int, error) {
	c, err := mdb.Collection(collectionTalks).CountDocuments(ctx, bson.D{})
	return int(c), wrapErr(err)
}

func (mdb *MongoDatabase) GetSimplePage(ctx context.Context, slug string) (res *models.SimplePage, err error) {
	filter := bson.D{
		{"slug", slug},
	}
	res = new(models.SimplePage)

	err = mdb.Collection(collectionPages).FindOne(ctx, filter).Decode(res)
	err = wrapErr(err)
	return
}

func wrapErr(err error) error {
	if err == nil {
		return err
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return fmt.Errorf("%w: %s", ErrNotFound, err.Error())
	}
	return fmt.Errorf("%w: %s", ErrGeneric, err.Error())
}
