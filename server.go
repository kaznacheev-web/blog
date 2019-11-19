package blog

import (
	"context"
	"github.com/kaznacheev-web/blog/internal/web"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/kaznacheev-web/blog/internal/config"
)

func Run(cfg config.MainConfig) error {
	r := mux.NewRouter()
	s := web.NewServer(cfg.Service, r)

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
