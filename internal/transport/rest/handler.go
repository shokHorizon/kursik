package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/shokHorizon/kursik/internal/transport/rest/handlers"
)

var (
	_ CoursesHandler = (*handlers.Courses)(nil)
	_ TasksHandler   = (*handlers.Tasks)(nil)
)

func ApplyCoursesRoutes(r chi.Router, h CoursesHandler) {
	// register courses routes
	r.Post("/course", h.CreateCourse)
	r.Get("/courses", h.GetCourses)
	r.Get("/course/{id}", h.GetCourse)
	r.Patch("/course", h.PatchCourse)
	r.Delete("/course/{id}", h.DeleteCourse)
}

func ApplyTasksRoutes(r chi.Router, h TasksHandler) {
	// register tasks routes
	r.Post("/task", h.CreateTask)
	r.Get("/tasks", h.GetTasks)
	r.Get("/task/{id}", h.GetTask)
	r.Patch("/task", h.PatchTask)
	r.Delete("/task/{id}", h.DeleteTask)
}
