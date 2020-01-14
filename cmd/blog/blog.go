package main

import (
	"flag"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kaznacheev-web/blog"
	"github.com/kaznacheev-web/blog/internal/config"
)

type args struct {
	ConfigPath string
}

func main() {
	a := args{}
	flag.StringVar(&a.ConfigPath, "-c", "configs/config.yml", "Path to configuration file")
	flag.Parse()

	cfg := config.MainConfig{}

	log.Printf("reading configuration from '%s'", a.ConfigPath)
	if err := cleanenv.ReadConfig(a.ConfigPath, &cfg); err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server at %s:%s", cfg.Service.Host, cfg.Service.Port)
	if err := blog.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
