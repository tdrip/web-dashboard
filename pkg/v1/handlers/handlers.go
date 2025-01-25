package handlers

import (
	"net/http"

	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type Router interface {
	// HTTP-method routing along `pattern`
	Connect(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Head(pattern string, h http.HandlerFunc)
	Options(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Trace(pattern string, h http.HandlerFunc)
}

type Handler struct {
	HttpMethod string
	Uri        string
}

type Handlers struct {
	Pages     []PageRenderHandler
	Partials  []PartialRenderHandler
	IPartials []PartialRenderHandler
}

func (h Handlers) RegisterRouter(router Router) {

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

func AddIPartial(h *Handlers, prh PartialRenderHandler) {
	handlers := h.IPartials
	handlers = append(handlers, prh)
	h.IPartials = handlers
}
