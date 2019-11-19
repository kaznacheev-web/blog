package web

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type DataProvider interface {
}

type Server struct {
	http.Server
	r *mux.Router
}

func NewServer(r *mux.Router) *Server {
	return &Server{
		Server: http.Server{
			Handler:      r,
			Addr:         "localhost:8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		r: r,
	}
}

func (s *Server) route() {

	api := s.r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/articles", s.handleArticlesGetAll())
	api.HandleFunc("/articles/{slug}", s.handleArticlesGetAll())
	api.HandleFunc("/talks", s.handleTalksGetAll())
	api.HandleFunc("/talks/{slug}", s.handleTalksGetOne())
	api.HandleFunc("/about", s.handleAboutGet())

	// s.r.HandleFunc("/", s.handleTemplate("index"))

	// r.GET("/", h.GetArticleList)
	// r.GET("/article/:slug", h.GetArticlePage)
	// r.GET("/talks", h.GetTalkList)
	// r.GET("/talks/:slug", h.GetTalkPage)
	// r.GET("/about", h.GetAboutPage)
}

func (s *Server) handleArticlesGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleArticlesGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleTalksGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleTalksGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleAboutGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleTemplate(templateName string, dataFunc func() interface{}) http.HandlerFunc {
	tpl, err := template.ParseFiles("static/templates/" + templateName)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.Execute(w, dataFunc()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
