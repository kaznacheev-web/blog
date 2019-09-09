package blog

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kaznacheev-web/blog/internal/config"
	"github.com/kaznacheev-web/blog/internal/database"
	"github.com/kaznacheev-web/blog/internal/handlers"
)

// StartServer runs website server with configuration
func StartServer(ctx context.Context, cfg config.MainConfig) error {
	r := gin.Default()

	db, err := database.NewMongoDatabase(cfg)
	if err != nil {
		return err
	}
	h := handlers.NewRequestHandler(ctx, cfg, db)

	r.GET("/", h.GetArticleList)
	r.GET("/article/:slug", h.GetArticlePage)
	r.GET("/talks", h.GetTalkList)
	r.GET("/talks/:slug", h.GetTalkPage)
	r.GET("/about", h.GetAboutPage)

	// admin := r.Group("/admin_box", gin.BasicAuth(gin.Accounts{}))

	// _ = admin

	log.Println(cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Service.Host, cfg.Service.Port),
		Handler: r,
	}

	// setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	select {
	case <-quit:
	case <-ctx.Done():
	}

	gsCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(gsCtx)
}

// Handler is an http request handler interface
type Handler interface {
	GetArticleList(*gin.Context)
	GetArticlePage(*gin.Context)
	GetTalkList(*gin.Context)
	GetTalkPage(*gin.Context)
	GetAboutPage(*gin.Context)
}
