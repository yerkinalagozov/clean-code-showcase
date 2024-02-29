package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/render"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/service"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	var user NewUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	newUserId, err := h.service.NewUser(r.Context(), user.MapToService())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]int{"id": newUserId})
}
