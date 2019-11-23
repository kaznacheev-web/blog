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

	GetTalks(page int) ([]models.Talk, error)
	GetTalk(slug string) (*models.Talk, error)
	GetTalkCount() (int, error)

	GetSimplePage(key string) (*models.SimplePage, error)
}

func (s *Server) HandleArticlesGetAll() http.HandlerFunc {
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

func (s *Server) HandleArticlesGetOne() http.HandlerFunc {
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

func (s *Server) HandleTalksGetAll() http.HandlerFunc {
	type response struct {
		Talks      []models.Talk
		TotalPages int
	}

	templateFunc := s.mustTemplate("talks")

	return func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.FormValue("page"))

		talks, err := s.sm.GetTalks(page)
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		total, err := s.sm.GetTalkCount()
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		d := response{
			Talks:      talks,
			TotalPages: total,
		}

		templateFunc(w, &d)
	}
}

func (s *Server) HandleTalksGetOne() http.HandlerFunc {
	templateFunc := s.mustTemplate("talk")

	return func(w http.ResponseWriter, r *http.Request) {
		slug := mux.Vars(r)["slug"]
		talk, err := s.sm.GetTalk(slug)
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		}

		templateFunc(w, talk)
	}
}

func (s *Server) HandleAboutGet() http.HandlerFunc {
	templateFunc := s.mustTemplate("about")

	return func(w http.ResponseWriter, r *http.Request) {
		page, err := s.sm.GetSimplePage("about")
		if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		}

		templateFunc(w, page)
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
