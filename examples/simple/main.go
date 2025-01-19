package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
	"github.com/tdrip/web-dashboard/pkg/v1/forms"
	navigation "github.com/tdrip/web-dashboard/pkg/v1/navigation"
	pages "github.com/tdrip/web-dashboard/pkg/v1/pages"
	render "github.com/tdrip/web-dashboard/pkg/v1/render"
	"github.com/tdrip/web-dashboard/pkg/v1/tables"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	router.Handle("/public/*", fileServer)

	pg := pages.Page{}
	pg.Title = "Test"
	pg.Id = "mainpage"
	pg.HasThemeSwicther = true
	pg.HeaderRawItem = bootstrap.DashBoardStyle
	pg.GetBodyMain = pg.GetEmptyMain
	//pg.BodyRawItem = `
	//<canvas class="my-4 w-100" id="myChart" width="1153" height="487" style="display: block; box-sizing: border-box; height: 487px; width: 1153px;"></canvas>
	//`
	pg.UseEmbeddedBootstrapCSS = false
	pg.UseEmbeddedBootstrapJS = false
	pg.UseEmbeddedDashBoardCSS = false
	pg.UseEmbeddedDashBoardJS = false
	pg.HeaderScripts = []string{
		"/public/color-modes.js",
	}

	pg.HeaderLinks = []pages.Link{
		{Item: "/public/bootstrap.min.css", Type: "stylesheet"},
		{Item: "/public/bootstrap-icons.min.css", Type: "stylesheet"},
		{Item: "/public/css@3.css", Type: "stylesheet"},
		{Item: "/public/dashboard.css", Type: "stylesheet"},
	}

	pg.BodyScripts = []string{
		"/public/bootstrap.bundle.min.js",
		//"/public/chart.umd.js",
		"/public/dashboard.js",
		"/public/htmgo.js",
	}

	pg.GetNavMenu = navigation.GetNavigation
	pg.GetBodyHeader = pg.SimpleNav
	pg.NavMenuItems = []navigation.NavMenuItem{
		/*
			{
				Title: "SHow table",
				Type:  2,
				HREF:  "/showtable",
			},
		*/
		navigation.NewNavButton("show table", "/showtable", "#mainpage"),
		/*{
			Title: "A Link",
			Type:  2,
			HREF:  "/",
		},
		{
			Title: "Test",
			Type:  1,
		},
		*/
	}

	//index
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderPage(request, writer, pg.GetPage)
	})
	sr := SimpleRender{}
	tableview := func(ctx *h.RequestContext) *h.Partial {
		return tables.RenderTable(sr)
	}

	router.Get("/showtable", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderPartial(request, writer, tableview)
	})

	ef := NewSillyForm()
	router.Get("/edititem/{id}", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, ef)
	})

	http.ListenAndServe(":3000", router)
}

type SillyForm struct {
	forms.EditForm
}

func NewSillyForm() SillyForm {
	sf := SillyForm{}
	sf.GetFormData = GetFakeFormData
	sf.Controls = []forms.FormControl{
		{
			Id:    "itemid",
			Title: "Item Title",
			Name:  "itemname",
			Cntrl: controls.TextInput{
				//BaseControlImp: controls.BaseControlImp{
				//	Text: "Ignored?",
				//},
				ReadOnly: false,
				Value:    "Unknown",
			},
		},
		{
			Id:    "itemselect",
			Title: "Select something",
			Name:  "itemname",
			Cntrl: controls.Select{
				Options: []controls.Option{
					controls.NewSimpleOption("opt1", "opt2"),
					controls.NewSimpleOption("opt2", "opt2"),
					controls.NewSimpleOption("opt3", "opt2"),
				},
			},
		},
	}
	return sf
}

func GetFakeFormData(c *h.RequestContext, ip render.IPartial) render.IPartial {
	cid := chi.URLParam(c.Request, "id")

	sf := ip.(forms.EditForm)
	sf.Title = fmt.Sprintf("editing %s", cid)

	cntrls := []forms.FormControl{}
	fmt.Printf("controls %d\n", len(cntrls))
	for _, cntl := range sf.Controls {
		ncntl := cntl
		if ncntl.Id == "itemid" {
			ct := cntl.Cntrl
			ti := ct.(controls.TextInput)
			ti.Value = fmt.Sprintf("Value: %s", cid)
			ti.ReadOnly = false
			//ti.BaseControlImp = controls.BaseControlImp{Text: "Updated?"}
			ncntl.Cntrl = ti
		}
		cntrls = append(cntrls, ncntl)
	}
	fmt.Printf("after controls %d\n", len(cntrls))
	sf.Controls = cntrls
	return sf
}

type SimpleRender struct {
	tables.TableRender
}

func (sr SimpleRender) GetHeaders() []string {
	return []string{"col1", "col2", "actions"}
}
func (sr SimpleRender) HasTitle() bool {
	return true
}
func (sr SimpleRender) GetTitle() string {
	return "Table title"
}

func (sr SimpleRender) HasNewButton() bool {
	return true
}
func (sr SimpleRender) HasUpdateTime() bool {
	return true
}
func (sr SimpleRender) GetModalCreateUrl() string {
	return "Table title"
}
func (sr SimpleRender) GetModalCreateId() string {
	return "Table title"
}
func (sr SimpleRender) GetModalEditUrl(string) string {
	return "Table title"
}
func (sr SimpleRender) GetTableBody() *h.Element {
	rows := []string{}

	for i := 0; i < 100; i++ {
		rows = append(rows, fmt.Sprintf("row%d", i))
	}

	return h.TBody(
		h.List(rows, MakeCells),
	)
}

func MakeCells(item string, index int) *h.Element {
	buttons := []controls.Button{
		{
			//BaseControlImp: controls.BaseControlImp{
			Text: "goo",
			Classes: []string{
				"btn",
				bootstrap.ButtonDanger,
			},
			Attributes: []*h.AttributeR{
				{
					Name:  "id",
					Value: fmt.Sprintf("testid%d", index),
				},
				{
					Name:  hx.TargetAttr,
					Value: "#mainpage",
				},
			},
			//},
			GetUrl: fmt.Sprintf(strings.Replace("/edititem/{id}", "{id}", fmt.Sprintf("testid%d", index), -1)),
		},
		{
			//BaseControlImp: controls.BaseControlImp{
			Text: "goo1",
			Classes: []string{
				"btn",
				bootstrap.ButtonSuccss,
			},
			//},
		},
	}
	return h.Tr(
		tables.GetTextCell(item+" col1"),
		tables.GetTextCell(item+" col2"),
		tables.GetButtonCell(buttons),
	)
}

func getCells(item string, index int) *h.Element {
	return tables.GetTextCell(item)
}
