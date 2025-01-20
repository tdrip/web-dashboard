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
		"/public/dashboard.js",
		"/public/htmgo.js",
	}

	pg.GetNavMenu = navigation.GetNavigation
	pg.GetBodyHeader = pg.SimpleNav
	pg.NavMenuItems = []navigation.NavMenuItem{
		navigation.NewNavButton("show table", "/showtable", "#mainpage"),
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
				ReadOnly: false,
				Value:    "Unknown",
			},
		},
		{
			Id:    "itemselect",
			Title: "Select something",
			Name:  "itemname",
			Cntrl: controls.Select{
				Classes: []string{
					bootstrap.FormSelect,
				},
				Options: []controls.Option{
					controls.NewSimpleOption("opt1", "opt2"),
					controls.NewSimpleOption("opt2", "opt2"),
					controls.NewSimpleOption("opt3", "opt2"),
				},
			},
		},
		{
			Id:    "itemtable",
			Title: "Update this table",
			Name:  "itemname",
			Cntrl: getTable(),
		},
	}
	sf.Buttons = []controls.Button{
		{
			Text: "Back",
			Classes: []string{
				bootstrap.Button,
				bootstrap.ButtonPrimary,
			},
			Attributes: []*h.AttributeR{
				{
					Name:  hx.TargetAttr,
					Value: "#mainpage",
				},
				{
					Name:  hx.GetAttr,
					Value: "/showtable",
				},
				{
					Name:  hx.SwapAttr,
					Value: hx.SwapTypeInnerHtml,
				},
			},
		},
		{
			Text: "Save",
			Classes: []string{
				bootstrap.Button,
				bootstrap.ButtonSuccess,
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

func getTable() controls.Table {
	tbl := controls.Table{}

	tbl.Classes = controls.SetClassses(tbl, []string{"table-responsive", "small", bootstrap.TableClass, "table-striped", "table-sm", "delete-row-example"})
	tbl.TableHeaders = controls.GetSimpleTableHeaders([]string{"col1", "col2", "actions"})
	tbl.TableBody = controls.TableBody{
		GetTableRows: GetTableRows,
	}
	return tbl
}

func (sr SimpleRender) GetTable() controls.Table {
	return getTable()
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

func GetTableRows() *h.Element {
	rows := []string{}

	for i := 0; i < 100; i++ {
		rows = append(rows, fmt.Sprintf("row%d", i))
	}
	return h.List(rows, MakeCells)
}

func MakeCells(item string, index int) *h.Element {
	buttons := []controls.Button{
		{

			Text: "goo",
			Classes: []string{
				bootstrap.Button,
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
			GetUrl: fmt.Sprintf(strings.Replace("/edititem/{id}", "{id}", fmt.Sprintf("testid%d", index), -1)),
		},
		{
			Text: "goo1",
			Classes: []string{
				bootstrap.Button,
				bootstrap.ButtonSuccess,
			},
		},
	}
	return h.Tr(
		controls.GetTextCell(item+" col1"),
		controls.GetTextCell(item+" col2"),
		controls.GetButtonCell(buttons),
	)
}

func getCells(item string, index int) *h.Element {
	return controls.GetTextCell(item)
}
