package homehandler

import (
	"net/http"

	"github.com/loissascha/go-http-server/server"
)

type HomeHandler struct {
}

func New() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) RegisterHandlers(s *server.Server) {
	s.GET("/", h.homeRoute)
}

func (h *HomeHandler) homeRoute(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
