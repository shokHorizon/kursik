package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/shokHorizon/kursik/internal/transport/rest/handlers"
	"net/http"
)

type Server struct {
	Router chi.Router
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) ApplyRoutes(courses handlers.Courses, tasks handlers.Tasks) {
	ApplyCoursesRoutes(s.Router, &courses)
	ApplyTasksRoutes(s.Router, &tasks)
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
