package web

// route sets up all routes and corresponding handlers
func (s *Server) route() {
	s.r.HandleFunc("/", s.HandleArticlesGetAll())
	s.r.HandleFunc("/article/{slug}", s.HandleArticlesGetOne())
	s.r.HandleFunc("/talks", s.HandleTalksGetOne())
	s.r.HandleFunc("/talks/{slug}", s.HandleTalksGetOne())
	s.r.HandleFunc("/about", s.HandleAboutGet())
}
