package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	bootstrap "github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	navigation "github.com/tdrip/web-dashboard/pkg/v1/navigation"
)

type GetNavMenu func(navitems []navigation.NavMenuItem) *h.Element
type GetModal func(id string) *h.Element
type GetBodyHeader func() *h.Element
type GetBodyMain func() *h.Element

type Link struct {
	Item string
	Type string
}

type Page struct {
	NavMenuItems            []navigation.NavMenuItem
	Header                  string
	Title                   string
	HeaderScripts           []string
	HeaderLinks             []Link
	HeaderRawItem           string
	BodyRawItem             string
	BodyScripts             []string
	GetNavMenu              GetNavMenu
	GetModal                GetModal
	GetBodyHeader           GetBodyHeader
	GetBodyMain             GetBodyMain
	HasThemeSwicther        bool
	UseEmbeddedBootstrapCSS bool
	UseEmbeddedBootstrapJS  bool
	UseEmbeddedDashBoardCSS bool
	UseEmbeddedDashBoardJS  bool
	UseEmbeddedChartUMDJS   bool
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
				pg.checkEmbeddedBootstrapCSS(),
				pg.checkEmbeddedDashboardCSS(),
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
				pg.checkGetNavMenu(),
				pg.checkGetBodyMain(),
			),
		),
		h.List(pg.BodyScripts, scriptItems),
		h.UnsafeRaw(pg.checkEmbeddedBootstrapJS()),
		h.UnsafeRaw(pg.checkEmbeddedChartUMDJS()),
		h.UnsafeRaw(pg.checkEmbeddedDashboardJS()),
	)
}

func (pg Page) checkEmbeddedChartUMDJS() string {

	if pg.UseEmbeddedDashBoardJS {
		return bootstrap.ChartUMD432JS
	}

	if !pg.UseEmbeddedChartUMDJS {
		return ""
	}
	return bootstrap.ChartUMD432JS
}

func (pg Page) checkEmbeddedDashboardJS() string {
	if !pg.UseEmbeddedDashBoardJS {
		return ""
	}
	return bootstrap.DashboardJS
}

func (pg Page) checkEmbeddedDashboardCSS() *h.Element {
	if !pg.UseEmbeddedDashBoardCSS {
		return h.Empty()
	}
	return h.Style(bootstrap.DashboardCSS)
}

func (pg Page) checkEmbeddedBootstrapCSS() *h.Element {
	if !pg.UseEmbeddedBootstrapCSS {
		return h.Empty()
	}
	return h.Style(bootstrap.CSS533)
}

func (pg Page) checkEmbeddedBootstrapJS() string {
	if !pg.UseEmbeddedBootstrapJS {
		return ""
	}
	return bootstrap.JS533
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
	if pg.GetModal == nil {
		return h.Empty()
	}
	return pg.GetModal()
}

func (pg Page) checkGetNavMenu() *h.Element {
	if pg.GetNavMenu == nil {
		return h.Empty()
	}
	return pg.GetNavMenu(pg.NavMenuItems)
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

func (pg Page) GetEmptyMain() *h.Element {
	return h.Main(
		h.Attribute("id", "page-data"),
		h.Class(bootstrap.ColMD9, "ms-sm-atuo", bootstrap.ColLG10, "px-md-4"),
		h.H2F(pg.Title),
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
