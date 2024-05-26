package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/usecase"
	"net/http"
	"strconv"
)

type Tasks struct {
	uc usecase.Tasks
}

func NewTasks(uc usecase.Tasks) *Tasks {
	return &Tasks{uc: uc}
}

func (h *Tasks) CreateTask(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		task entity.Task
		err  error
	)
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.uc.CreateTask(ctx, &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Tasks) GetTasks(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        = r.Context()
		pagination entity.Pagination
		filter     entity.TaskFilter
		tasks      entity.Tasks
		err        error
	)
	err = json.NewDecoder(r.Body).Decode(&pagination)
	err = json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tasks, err = h.uc.GetTasks(ctx, filter, pagination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Tasks) GetTask(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		id   int
		task *entity.Task
		err  error
	)
	id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err = h.uc.GetTask(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Tasks) PatchTask(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		task entity.Task
		err  error
	)
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.uc.PatchTask(ctx, &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Tasks) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  int
		err error
	)
	id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.uc.DeleteTask(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
