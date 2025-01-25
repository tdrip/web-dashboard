package render

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

type Link struct {
	Item string
	Type string
}

type Site struct {
	Title         string
	HeaderScripts []string
	HeaderLinks   []Link
	HeaderRawItem string
	PageBody      controls.PageBody
}

func (ctrl Site) GetPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		h.Html(
			h.HxExtensions(
				h.BaseExtensions(),
			),
			h.Attribute("lang", "en"),
			h.Attribute("data-bs-theme", "light"),
			h.Head(
				h.Meta("viewport", "wid=device-width initial-scale=1"),
				h.Meta("theme-color", "#712cf9"),
				h.Title(h.Text(ctrl.Title)),
				h.List(ctrl.HeaderScripts, scriptItems),
				getLinks(ctrl.HeaderLinks),
				h.UnsafeRaw(ctrl.HeaderRawItem),
			),
			ctrl.PageBody.ToHTML(),
		),
	)
}

func getLinks(links []Link) *h.Element {
	return h.List(links, linkItems)
}

func linkItems(item Link, index int) *h.Element {
	return h.Link(item.Item, item.Type)
}

func scriptItems(item string, index int) *h.Element {
	return h.Script(item)
}
