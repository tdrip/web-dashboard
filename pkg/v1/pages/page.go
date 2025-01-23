package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	bootstrap "github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

type GetModal func() *h.Element
type GetBodyHeader func() *h.Element
type GetBodyMain func() *h.Element

type Link struct {
	Item string
	Type string
}

type Page struct {
	SidebarMenu      controls.SidebarMenu
	Header           string
	Title            string
	HeaderScripts    []string
	HeaderLinks      []Link
	HeaderRawItem    string
	BodyRawItem      string
	BodyScripts      []string
	GetModal         GetModal
	HasModal         bool
	ModalID          string
	GetBodyHeader    GetBodyHeader
	GetBodyMain      GetBodyMain
	HasThemeSwicther bool
	Id               string
}

func (pg Page) GetPage(ctx *h.RequestContext) *h.Page {
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
				h.Title(h.Text(pg.Title)),
				h.List(pg.HeaderScripts, scriptItems),
				getLinks(pg.HeaderLinks),
				h.UnsafeRaw(pg.HeaderRawItem),
			),
			pg.checkPage(),
		),
	)
}

func (pg Page) checkPage() *h.Element {
	return h.Body(
		h.UnsafeRaw(pg.BodyRawItem),
		h.UnsafeRaw(pg.checkGetThemeSwitcher()),
		pg.checkGetBodyHeader(),
		pg.checkGetModals(),
		h.Div(
			h.Class(bootstrap.ContainerFluid),
			h.Div(
				h.Class(bootstrap.Row),
				pg.checkGetSidebarMenu(),
				h.Main(
					h.Class(bootstrap.ColMD9, "ms-sm-atuo", bootstrap.ColLG10, "px-md-4"),
					pg.GetPageTitle(),
					//h.Div(
					h.Attribute("id", pg.Id),
					pg.checkGetBodyMain(),
					//),
				),
			),
		),
		h.List(pg.BodyScripts, scriptItems),
	)
}

func (pg Page) checkGetThemeSwitcher() string {
	if !pg.HasThemeSwicther {
		return ""
	}
	return bootstrap.ThemeSwitcher
}

func (pg Page) checkGetBodyMain() *h.Element {
	if pg.GetBodyMain == nil {
		return h.Empty()
	}
	return pg.GetBodyMain()
}

func (pg Page) checkGetBodyHeader() *h.Element {
	if pg.GetBodyHeader == nil {
		return h.Empty()
	}
	return pg.GetBodyHeader()
}

func (pg Page) checkGetModals() *h.Element {
	if !pg.HasModal {
		return h.Empty()
	}
	return pg.GetModal()
}

func (pg Page) GetEmptyModal() *h.Element {
	return h.Div(
		h.Attribute("id", pg.ModalID),
		h.Class(bootstrap.Modal, bootstrap.ModalBlur, "fade"),
		h.Attribute("style", "display:none"),
		h.AriaHidden(false),
		h.TabIndex(-11),
		h.Div(
			h.Class(bootstrap.ModalDialog, bootstrap.ModalLG, bootstrap.ModalDialogCentered),
			h.Attribute("role", "document"),
			h.Div(
				h.Class(bootstrap.ModalContent),
			),
		),
	)
}

func (pg Page) checkGetSidebarMenu() *h.Element {
	if len(pg.SidebarMenu.MenuItems) == 0 {
		return h.Empty()
	}
	return pg.SidebarMenu.ToHTML()
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

func (pg Page) GetPageTitle() *h.Element {
	return h.Div(
		h.Class("d-flex", "justify-content-between", "flex-wrap", "flex-md-nowrap", "align-items-center", "pt-3", "pb-2", "mb-3", "border-bottom"),
		h.H1F(pg.Title),
	)
}

func (pg Page) SimpleNav() *h.Element {
	return h.Header(
		h.Class(bootstrap.NavBarClass, bootstrap.NavBarDarkClass, bootstrap.StickyTop, bootstrap.BGDark, bootstrap.FlexMDNoWrap, "p-0", "shadow"),
		h.A(
			h.Class(bootstrap.NavBarBrandClass, bootstrap.ColMD9, bootstrap.ColLG2, "me-0", "px-3", "fs-6"),
			h.Href("/"),
			h.Text(pg.Title),
		),
	)
}
