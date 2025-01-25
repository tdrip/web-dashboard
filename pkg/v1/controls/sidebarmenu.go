package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type SidebarMenu struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	MenuItems  []NavMenuItem
}

func (ctrl SidebarMenu) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl SidebarMenu) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl SidebarMenu) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl SidebarMenu) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl SidebarMenu) ToHTML() *h.Element {
	return h.Div(
		//h.Attribute("id", "sidebarMenu"),
		h.Class("sidebar", "border", "border-right", bootstrap.ColMD3, bootstrap.ColLG2, "p-0", "bg-body-tertiary"),
		// "collapse" "d-md-block", ,
		h.Div(
			h.Attribute(bootstrap.AttributeId, "sidebarMenu"),
			//h.Class("position-sticky", "pt-3", bootstrap.StickySidebar),
			h.Class("offcanvas-md", "offcanvas-end", "bg-body-tertiary"),
			h.Div(
				h.Class("offcanvas-body", "d-md-flex", bootstrap.FlexColumn, "p-0", "pt-lg-3", "overflow-y-auto"),
				getNavList(ctrl.MenuItems),
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
	case Heading:
		return h.H6(
			h.Class("sidebar-heading", "d-flex", "justify-content-between", "align-items-center", "px-3", "mt4", "mb-1", "text-body-secondary", "text-uppercase"),
			h.Span(h.Text(item.Title)),
		)
	case NavHrefBlank:
		return h.Li(
			h.Class(bootstrap.NavItem),
			h.A(
				h.Class("nav-link", "d-flex", "align-items-center", "gap-2"),
				h.Attribute("href", item.HREF),
				h.Attribute("target", "_blank"),
				h.Text(item.Title),
			),
		)
	case NavHref:
		return h.Li(
			h.Class(bootstrap.NavItem),
			h.A(
				h.Class("nav-link", "d-flex", "align-items-center", "gap-2"),
				h.Attribute("href", item.HREF),
				h.Text(item.Title),
			),
		)
	case NavButton:
		return h.Li(
			h.Class(bootstrap.NavItem),
			h.Button(
				h.Class("nav-link", "active", "d-flex", "align-items-center", "gap-2"),
				h.Attribute("aria-current", "page"),
				h.HxTarget(item.Target), //"#page-data"),
				h.Attribute(hx.GetAttr, item.HREF),
				h.Attribute(hx.SwapAttr, hx.SwapTypeInnerHtml),
				h.Text(item.Title),
			),
		)

	case NavHRule:
		return h.Hr(
			h.Class("my-3"),
		)
	}
	return h.Li(h.Text(item.Title))
}
