package homehandler

import (
	"net/http"

	"github.com/loissascha/go-http-server/server"
	"github.com/loissascha/go-templ-template/internal/ui/components"
	"github.com/loissascha/go-templ-template/internal/ui/layouts"
	"github.com/loissascha/go-templ-template/internal/ui/pages"
)

type HomeHandler struct {
}

func New() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) RegisterHandlers(s *server.Server) {
	s.GET("/", h.homeRoute)
	s.GET("/lorem", h.loremIpsumRoute)
}

func (h *HomeHandler) homeRoute(w http.ResponseWriter, r *http.Request) {
	homeComponent := pages.Home()
	layoutComponent := layouts.Layout(homeComponent)
	layoutComponent.Render(r.Context(), w)
}

func (h *HomeHandler) loremIpsumRoute(w http.ResponseWriter, r *http.Request) {
	components.TestInfo().Render(r.Context(), w)
}
