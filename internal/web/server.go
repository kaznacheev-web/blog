package web

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/kaznacheev-web/blog/internal/config"
	// "github.com/nicksnyder/go-i18n/v2/i18n"
)

type templateHandlerFunc func(http.ResponseWriter, interface{})

type Server struct {
	http.Server
	r                     *mux.Router
	sm                    StorageManager
	rootPath              string
	respErrorFunc         templateHandlerFunc
	respErrorNotFoundFunc templateHandlerFunc
}

func NewServer(cfg config.Server, r *mux.Router, sm StorageManager, rootPath string) *Server {
	s := &Server{
		Server: http.Server{
			Handler:      r,
			Addr:         cfg.Host + ":" + cfg.Port,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		r:        r,
		sm:       sm,
		rootPath: rootPath,
	}
	s.respErrorFunc = s.mustTemplate("error")
	s.respErrorNotFoundFunc = s.mustTemplate("error404")

	s.route()
	return s
}

func (s *Server) mustTemplate(templateName string) templateHandlerFunc {
	tpl, err := template.ParseFiles(s.rootPath + "/static/templates/" + templateName)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, data interface{}) {
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
