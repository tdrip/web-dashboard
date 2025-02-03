package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	bootstrap "github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type PageBody struct {
	BaseControl
	Attributes       []*h.AttributeR
	Classes          []string
	SidebarMenu      SidebarMenu
	Title            string
	RawItem          string
	Scripts          []string
	GetModal         DrawControl
	HasModal         bool
	GetBodyHeader    DrawControl
	GetBodyMain      DrawControl
	HasThemeSwicther bool
	Id               string
	BodyId           string
	HeaderId         string
	ModalId          string
}

func (ctrl PageBody) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl PageBody) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl PageBody) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl PageBody) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl PageBody) ToHTML() *h.Element {
	return h.Body(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		h.UnsafeRaw(ctrl.RawItem),
		h.UnsafeRaw(ctrl.checkGetThemeSwitcher()),
		ctrl.checkGetBodyHeader(),
		ctrl.checkGetModals(),
		h.Div(
			h.Class(bootstrap.ContainerFluid),
			h.Div(
				h.Class(bootstrap.Row),
				ctrl.checkGetSidebarMenu(),
				h.Main(
					h.Class(bootstrap.ColMD9, "ms-sm-atuo", bootstrap.ColLG10, "px-md-4"),
					ctrl.GetPageTitle(),
					h.Attribute("id", ctrl.Id),
					ctrl.checkGetBodyMain(),
				),
			),
		),
		h.List(ctrl.Scripts, scriptItems),
	)
}

func (ctrl PageBody) checkGetThemeSwitcher() string {
	if !ctrl.HasThemeSwicther {
		return ""
	}
	return bootstrap.ThemeSwitcher
}

func (ctrl PageBody) checkGetBodyMain() *h.Element {
	if ctrl.GetBodyMain == nil {
		return h.Empty()
	}
	return ctrl.GetBodyMain(ctrl.BodyId)
}

func (ctrl PageBody) checkGetBodyHeader() *h.Element {
	if ctrl.GetBodyHeader == nil {
		return h.Empty()
	}
	return ctrl.GetBodyHeader(ctrl.HeaderId)
}

func (ctrl PageBody) checkGetModals() *h.Element {
	if !ctrl.HasModal {
		return h.Empty()
	}
	return ctrl.GetModal(ctrl.ModalId)
}

func (ctrl PageBody) checkGetSidebarMenu() *h.Element {
	if len(ctrl.SidebarMenu.MenuItems) == 0 {
		return h.Empty()
	}
	return ctrl.SidebarMenu.ToHTML()
}

func scriptItems(item string, index int) *h.Element {
	return h.Script(item)
}

func (ctrl PageBody) GetPageTitle() *h.Element {
	return h.Div(
		h.Class("d-flex", "justify-content-between", "flex-wrap", "flex-md-nowrap", "align-items-center", "pt-3", "pb-2", "mb-3", "border-bottom"),
		h.H1F(ctrl.Title),
	)
}

func (ctrl PageBody) SimpleNav(Id string) *h.Element {
	return h.Header(
		h.Class(bootstrap.NavBarClass, bootstrap.NavBarDarkClass, bootstrap.StickyTop, bootstrap.BGDark, bootstrap.FlexMDNoWrap, "p-0", "shadow"),
		h.A(
			h.Class(bootstrap.NavBarBrandClass, bootstrap.ColMD9, bootstrap.ColLG2, "me-0", "px-3", "fs-6"),
			h.Href("/"),
			h.Text(ctrl.Title),
		),
	)
}
