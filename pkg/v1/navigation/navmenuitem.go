package navigation

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type NavMenuItem struct {
	Type  int
	Title string
	HREF  string
}

func GetNavigation(navitems []NavMenuItem) *h.Element {
	return h.Div(
		//h.Attribute("id", "sidebarMenu"),
		h.Class("sidebar", "border", "border-right", bootstrap.ColMD3, bootstrap.ColLG2, "p-0", "bg-body-tertiary"),
		// "collapse" "d-md-block", ,
		h.Div(
			h.Attribute("id", "sidebarMenu"),
			//h.Class("position-sticky", "pt-3", bootstrap.StickySidebar),
			h.Class("offcanvas-md", "offcanvas-end", "bg-body-tertiary"),
			h.Div(
				h.Class("offcanvas-body", "d-md-flex", bootstrap.FlexColumn, "p-0", "pt-lg-3", "overflow-y-auto"),
				getNavList(navitems),
			),
		),
	)
}

func getNavList(navitems []NavMenuItem) *h.Element {
	return h.Ul(
		h.Class(bootstrap.NavClass, bootstrap.FlexColumn),
		h.List(navitems, renderListItems),
	)
}

func renderListItems(item NavMenuItem, index int) *h.Element {

	switch item.Type {
	case 1:
		return h.H6(
			h.Class("sidebar-heading", "d-flex", "justify-content-between", "align-items-center", "px-3", "mt4", "mb-1", "text-body-secondary", "text-uppercase"),
			h.Span(h.Text(item.Title)),
		)
	case 2:
		return h.Li(
			h.Class("nav-item"),
			h.A(
				h.Class("nav-link"),
				h.Attribute("href", item.HREF),
				h.Attribute("target", "_blank"),
				h.Text(item.Title),
			),
		)
	case 3:
		return h.Li(
			h.Class("nav-item"),
			h.Button(
				h.Class("nav-link", "active"),
				h.Attribute("aria-current", "page"),
				h.HxTarget("#page-data"),
				h.Attribute(hx.GetAttr, item.HREF),
				h.Attribute(hx.SwapAttr, hx.SwapTypeInnerHtml),
				h.Text(item.Title),
			),
		)
	}
	return h.Li(h.Text(item.Title))
}
