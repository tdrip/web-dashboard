package handlers

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/modals"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
	"github.com/tdrip/web-dashboard/pkg/v1/tables"
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

func NewModalRenderHandler(method string, uri string, mr modals.ModalRender) PartialRenderHandler {
	partial := func(ctx *h.RequestContext) *h.Partial {
		return modals.RenderModal(mr)
	}
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: partial}
}

func NewTableRenderHandler(method string, uri string, tr tables.TableRender) PartialRenderHandler {
	partial := func(ctx *h.RequestContext) *h.Partial {
		return tables.RenderTable(tr)
	}
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: partial}
}

func NewPartialRenderHandler(method string, uri string, exec GetRender) PartialRenderHandler {
	return PartialRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}

func NewPageRenderHandler(method string, uri string, exec GetFullRender) PageRenderHandler {
	return PageRenderHandler{Handler: Handler{HttpMethod: method, Uri: uri}, Renderer: exec}
}
