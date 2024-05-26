package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/shokHorizon/kursik/internal/usecase"
	"net/http"
)

type Server struct {
	Handler http.Handler
	Router  chi.Router
}

func NewServer(h http.Handler) *Server {
	return &Server{
		Handler: h,
		Router:  chi.NewRouter(),
	}
}

func (s *Server) ApplyRoutes(courses usecase.Courses, tasks usecase.Tasks) {
	// register courses routes
	s.Router.Post("/course", courses.Create)
	s.Router.Get("/courses", courses.GetAll)
	s.Router.Get("/course/{id}", courses.Get)
	s.Router.Patch("/course", courses.Patch)
	s.Router.Delete("/course/{id}", courses.Delete)
	// register tasks routes
	s.Router.Post("/task", tasks.Create)
	s.Router.Get("/tasks", tasks.GetAll)
	s.Router.Get("/task/{id}", tasks.Get)
	s.Router.Patch("/task", tasks.Patch)
	s.Router.Delete("/task/{id}", tasks.Delete)
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Handler)
}
