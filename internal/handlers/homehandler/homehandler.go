package homehandler

import (
	"net/http"

	"github.com/loissascha/go-http-server/server"
	"github.com/loissascha/go-templ-template/internal/ui/components"
	"github.com/loissascha/go-templ-template/internal/ui/layouts"
	"github.com/loissascha/go-templ-template/internal/ui/pages"
)

type HomeHandler struct {
	s *server.Server
}

func New(s *server.Server) *HomeHandler {
	return &HomeHandler{
		s: s,
	}
}

func (h *HomeHandler) RegisterHandlers(s *server.Server) {
	s.GET("/", h.homeRoute)
	s.GET("/lorem", h.loremIpsumRoute)
}

func (h *HomeHandler) homeRoute(w http.ResponseWriter, r *http.Request) {
	lang := h.s.GetActiveLanguage(r)
	t := h.s.GetLanguageStringMap(r)
	homeComponent := pages.Home(t, lang)
	layoutComponent := layouts.Layout(homeComponent)
	layoutComponent.Render(r.Context(), w)
}

func (h *HomeHandler) loremIpsumRoute(w http.ResponseWriter, r *http.Request) {
	t := h.s.GetLanguageStringMap(r)
	components.TestInfo(t).Render(r.Context(), w)
}
