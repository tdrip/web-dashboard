package handlers

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type GetRender func(c *h.RequestContext) *h.Partial
type GetFullRender func(c *h.RequestContext) *h.Page

type PartialRenderHandler struct {
	Handler
	Renderer GetRender
	IPartial render.IPartial
}

type PageRenderHandler struct {
	Handler
	Renderer GetFullRender
}

func NewIPartialRenderHandler(method string, uri string, render render.IPartial) PartialRenderHandler {
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, IPartial: render}
}

func NewPartialRenderHandler(method string, uri string, exec GetRender) PartialRenderHandler {
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}

func NewPageRenderHandler(method string, uri string, exec GetFullRender) PageRenderHandler {
	return PageRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}
