package web

// route sets up all routes and corresponding handlers
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
