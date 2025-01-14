package handlers

import "github.com/maddalax/htmgo/framework/h"

type GetRender func(c *h.RequestContext) *h.Partial
type GetFullRender func(c *h.RequestContext) *h.Page

type PartialRenderHandler struct {
	Handler
	Renderer GetRender
}

type PageRenderHandler struct {
	Handler
	Renderer GetFullRender
}

func NewPartialRenderHandler(method string, uri string, exec GetRender) PartialRenderHandler {
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}

func NewPageRenderHandler(method string, uri string, exec GetFullRender) PageRenderHandler {
	return PageRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}
