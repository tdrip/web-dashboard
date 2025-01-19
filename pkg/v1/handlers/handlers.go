package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type Handler struct {
	HttpMethod string
	Uri        string
}

type Handlers struct {
	Pages     []PageRenderHandler
	Partials  []PartialRenderHandler
	IPartials []PartialRenderHandler
}

func (h Handlers) RegisterRouter(router chi.Router) {

	for _, hndlr := range h.IPartials {
		switch hndlr.HttpMethod {
		case http.MethodGet:
			router.Get(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderIPartial(request, writer, hndlr.IPartial)
			})
		case http.MethodDelete:
			router.Delete(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderIPartial(request, writer, hndlr.IPartial)
			})
		case http.MethodPost:
			router.Post(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderIPartial(request, writer, hndlr.IPartial)
			})
		case http.MethodPut:
			router.Put(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderIPartial(request, writer, hndlr.IPartial)
			})
		case http.MethodHead:
			router.Head(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderIPartial(request, writer, hndlr.IPartial)
			})
		}
	}

	for _, hndlr := range h.Partials {
		switch hndlr.HttpMethod {
		case http.MethodGet:
			router.Get(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPartial(request, writer, hndlr.Renderer)
			})
		case http.MethodDelete:
			router.Delete(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPartial(request, writer, hndlr.Renderer)
			})
		case http.MethodPost:
			router.Post(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPartial(request, writer, hndlr.Renderer)
			})
		case http.MethodPut:
			router.Put(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPartial(request, writer, hndlr.Renderer)
			})
		case http.MethodHead:
			router.Head(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPartial(request, writer, hndlr.Renderer)
			})
		}
	}

	for _, hndlr := range h.Pages {
		switch hndlr.HttpMethod {
		case http.MethodGet:
			router.Get(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPage(request, writer, hndlr.Renderer)
			})
		case http.MethodDelete:
			router.Delete(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPage(request, writer, hndlr.Renderer)
			})
		case http.MethodPost:
			router.Post(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPage(request, writer, hndlr.Renderer)
			})
		case http.MethodPut:
			router.Put(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPage(request, writer, hndlr.Renderer)
			})
		case http.MethodHead:
			router.Head(hndlr.Uri, func(writer http.ResponseWriter, request *http.Request) {
				render.RenderPage(request, writer, hndlr.Renderer)
			})
		}
	}
}

func AddPage(h *Handlers, prh PageRenderHandler) {
	handlers := h.Pages
	handlers = append(handlers, prh)
	h.Pages = handlers
}

func AddPartial(h *Handlers, prh PartialRenderHandler) {
	handlers := h.Partials
	handlers = append(handlers, prh)
	h.Partials = handlers
}
