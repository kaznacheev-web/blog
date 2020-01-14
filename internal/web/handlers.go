package web

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kaznacheev-web/blog/internal/database"
	"github.com/kaznacheev-web/blog/internal/models"
)

//go:generate mockery -name=StorageManager -output=mocks

// StorageManager processes storage operations
type StorageManager interface {
	// Read operations

	// GetArticles returns a list of articles limited by the page number
	GetArticles(ctx context.Context, page int) ([]models.Article, int, error)
	// GetArticle returns a certain article
	GetArticle(ctx context.Context, slug string) (*models.Article, error)
	GetArticleCount(ctx context.Context) (int, error)
	// GetTalks returns a list of talks limited by the page number
	GetTalks(ctx context.Context, page int) ([]models.Talk, int, error)
	GetTalk(ctx context.Context, slug string) (*models.Talk, error)
	GetTalkCount(ctx context.Context) (int, error)

	GetSimplePage(ctx context.Context, slug string) (*models.SimplePage, error)

	// Update operations

	// IncrementArticleVCounter(slug string) error
	// IncrementTalkVCounter(slug string) error
}

func (s *Server) HandleArticlesGetAll() http.HandlerFunc {
	type response struct {
		Articles   []models.Article
		TotalPages int
	}

	// templateFunc := s.mustTemplate("articles")

	return func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.FormValue("page"))

		articles, total, err := s.sm.GetArticles(r.Context(), page)
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

		// templateFunc(w, &d)
		json.NewEncoder(w).Encode(&d)
	}
}

func (s *Server) HandleArticlesGetOne() http.HandlerFunc {
	// templateFunc := s.mustTemplate("article")

	return func(w http.ResponseWriter, r *http.Request) {
		slug := mux.Vars(r)["slug"]

		article, err := s.sm.GetArticle(r.Context(), slug)
		if errors.Is(err, database.ErrNotFound) {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		} else if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		// templateFunc(w, article)
		json.NewEncoder(w).Encode(article)
	}
}

func (s *Server) HandleTalksGetAll() http.HandlerFunc {
	type response struct {
		Talks      []models.Talk
		TotalPages int
	}

	// templateFunc := s.mustTemplate("talks")

	return func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.FormValue("page"))

		talks, total, err := s.sm.GetTalks(r.Context(), page)
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

		// templateFunc(w, &d)
		json.NewEncoder(w).Encode(&d)
	}
}

func (s *Server) HandleTalksGetOne() http.HandlerFunc {
	// templateFunc := s.mustTemplate("talk")

	return func(w http.ResponseWriter, r *http.Request) {
		slug := mux.Vars(r)["slug"]
		talk, err := s.sm.GetTalk(r.Context(), slug)
		if errors.Is(err, database.ErrNotFound) {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		} else if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		// templateFunc(w, talk)
		json.NewEncoder(w).Encode(talk)
	}
}

func (s *Server) HandleAboutGet() http.HandlerFunc {
	// templateFunc := s.mustTemplate("about")

	return func(w http.ResponseWriter, r *http.Request) {
		page, err := s.sm.GetSimplePage(r.Context(), "about")
		if errors.Is(err, database.ErrNotFound) {
			s.handleError(w, errorMessage{
				Status: http.StatusNotFound,
				Text:   err.Error(),
			})
			return
		} else if err != nil {
			s.handleError(w, errorMessage{
				Status: http.StatusInternalServerError,
				Text:   err.Error(),
			})
			return
		}

		// templateFunc(w, page)
		json.NewEncoder(w).Encode(page)
	}
}

type errorMessage struct {
	Status int
	Text   string
}

func (s *Server) handleError(w http.ResponseWriter, msg errorMessage) {
	log.Printf("error %d: %s", msg.Status, msg.Text)
	switch msg.Status {
	case http.StatusNotFound:
		s.respErrorNotFoundFunc(w, msg)
	default:
		s.respErrorFunc(w, msg)
	}
}
