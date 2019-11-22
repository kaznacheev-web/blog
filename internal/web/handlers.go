package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kaznacheev-web/blog/internal/models"
)

// StorageManager processes storage operations
type StorageManager interface {
	GetArticles(page int) ([]models.Article, error)
	GetArticle(slug string) (*models.Article, error)
	GetArticleCount() (int, error)
}

func (s *Server) handleArticlesGetAll() http.HandlerFunc {
	type response struct {
		Articles   []models.Article
		TotalPages int
	}

	templateFunc := s.mustTemplate("articles")

	return func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.FormValue("page"))

		articles, err := s.sm.GetArticles(page)
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		total, err := s.sm.GetArticleCount()
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		d := response{
			Articles:   articles,
			TotalPages: total,
		}

		templateFunc(w, &d)
	}
}

func (s *Server) handleArticlesGetOne() http.HandlerFunc {
	templateFunc := s.mustTemplate("article")

	return func(w http.ResponseWriter, r *http.Request) {
		slug := mux.Vars(r)["slug"]
		article, err := s.sm.GetArticle(slug)
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		}

		templateFunc(w, article)
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

type errorMessage struct {
	Status int
	Text   string
}

func (s *Server) handleError(w http.ResponseWriter, msg errorMessage) {
	switch msg.Status {
	case http.StatusNotFound:
		s.respErrorNotFoundFunc(w, &msg)
	default:
		s.respErrorFunc(w, &msg)
	}
}
