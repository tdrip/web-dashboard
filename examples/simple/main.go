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
	//pg.GetBodyMain = pg.GetEmptyMain
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
		navigation.NewNavButton("show grid form", "/gridform", "#mainpage"),
		navigation.NewNavHRule(),
	}

	//index
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderPage(request, writer, pg.GetPage)
	})

	router.Get("/gridform", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, NewGF())
	})

	router.Get("/edititem/{id}", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, NewSillyForm())
	})

	http.ListenAndServe(":3000", router)
}

func NewGF() forms.GridForm {
	gf := forms.GridForm{}
	gf.Title = "A grid form"
	gf.HasUpdateTime = true
	gf.GetTable = getTable

	newbtn := controls.Button{
		Text: "New",
		Classes: []string{
			bootstrap.Button,
			bootstrap.ButtonSuccess,
		},
		GetUrl: "/getmodal/",
		Attributes: []*h.AttributeR{
			{
				Name:  hx.TargetAttr,
				Value: "#topmodal",
			},
			{
				Name:  "data-bs-toggle",
				Value: "modal",
			},
			{
				Name:  "data-bs-target",
				Value: "#topmodal",
			},
		},
	}
	gf.NewButton = &newbtn
	return gf
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
				ReadOnly:   false,
				Attributes: []*h.AttributeR{controls.GetAttValue("Unknown")},
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
			ti.Attributes = controls.SetAttValue(ti, fmt.Sprintf("Value: %s", cid))
			ti.ReadOnly = false
			ncntl.Cntrl = ti
		}
		cntrls = append(cntrls, ncntl)
	}
	fmt.Printf("after controls %d\n", len(cntrls))
	sf.Controls = cntrls
	return sf
}

func getTable() controls.Table {
	tbl := controls.Table{}

	tbl.Classes = controls.SetClasses(tbl, []string{"table-responsive", "small", bootstrap.TableClass, "table-striped", "table-sm", "delete-row-example"})
	tbl.TableHeaders = controls.GetSimpleTableHeaders([]string{"unused?", "col1", "col2", "actions"})
	tbl.TableBody = controls.TableBody{
		GetTableRows: GetTableRows,
	}
	return tbl
}

func GetTableRows() *h.Element {
	rows := []string{}

	for i := 0; i < 100; i++ {
		rows = append(rows, fmt.Sprintf("row%d", i))
	}
	return h.List(rows, MakeCells)
}

func MakeCells(item string, index int) *h.Element {

	id := fmt.Sprintf("testid%d", index)

	cbox := controls.NewCheckedCheckbox(id+" displayname", id+"-chk", id+"-chk-name")

	if index%2 == 0 {
		cbox = controls.NewUnCheckedCheckbox(id+" displayname", id+"-chk", id+"-chk-name")
	}
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
					Value: id,
				},
				{
					Name:  hx.TargetAttr,
					Value: "#mainpage",
				},
			},
			GetUrl: strings.Replace("/edititem/{id}", "{id}", id, -1),
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
		controls.GetCheckBoxCell(cbox),
		controls.GetTextCell(item+" col1"),
		controls.GetTextCell(item+" col2"),
		controls.GetButtonCell(buttons),
	)
}
