package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/usecase"
	"net/http"
	"strconv"
)

type Courses struct {
	uc usecase.Courses
}

func NewCourses(uc usecase.Courses) *Courses {
	return &Courses{uc: uc}
}

func (h *Courses) CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var course entity.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.uc.Create(ctx, &course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Courses) GetCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var pagination entity.Pagination
	err := json.NewDecoder(r.Body).Decode(&pagination)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	courses, err := h.uc.GetAll(ctx, pagination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Courses) GetCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	course, err := h.uc.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Courses) PatchCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var course entity.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.uc.Patch(ctx, &course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Courses) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.uc.Delete(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
