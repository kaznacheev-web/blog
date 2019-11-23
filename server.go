package blog

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/kaznacheev-web/blog/internal/config"
	"github.com/kaznacheev-web/blog/internal/database"
	"github.com/kaznacheev-web/blog/internal/web"
)

func Run(cfg config.MainConfig) error {
	r := mux.NewRouter()

	mdb, err := database.NewMongoDatabase(cfg.Database)
	if err != nil {
		return err
	}

	s := web.NewServer(cfg.Service, r, mdb, cfg.RootDir)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("starting server")
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("shutting down")
	return s.Shutdown(ctx)
}
