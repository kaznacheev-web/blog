package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kaznacheev-web/blog/internal/config"
	// "github.com/nicksnyder/go-i18n/v2/i18n"
)

// RequestHandler is an http request handler
type RequestHandler struct {
	db  Database
	ctx context.Context
}

// NewRequestHandler creates a new http request handler
func NewRequestHandler(ctx context.Context, cfg config.MainConfig, db Database) *RequestHandler {
	return &RequestHandler{
		db:  db,
		ctx: ctx,
	}
}

// GetArticleList returns a list of articles
func (h RequestHandler) GetArticleList(c *gin.Context) {
	// lang := r.FormValue("lang")
	// accept := r.Header.Get("Accept-Language")
	// localizer := i18n.NewLocalizer(bundle, lang, accept)
}

func (h RequestHandler) GetArticlePage(c *gin.Context) {

}

func (h RequestHandler) GetTalkList(c *gin.Context) {

}

func (h RequestHandler) GetTalkPage(c *gin.Context) {

}

func (h RequestHandler) GetAboutPage(c *gin.Context) {

}

// Database is a set of database operations
type Database interface {
}
