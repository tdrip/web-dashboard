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
	render "github.com/tdrip/web-dashboard/pkg/v1/render"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	router.Handle("/public/*", fileServer)

	site := render.Site{}
	site.Title = "Test"
	site.HeaderRawItem = bootstrap.DashBoardStyle
	site.HeaderScripts = []string{
		"/public/color-modes.js",
	}

	site.HeaderLinks = []render.Link{
		{Item: "/public/bootstrap.min.css", Type: "stylesheet"},
		{Item: "/public/bootstrap-icons.min.css", Type: "stylesheet"},
		{Item: "/public/css@3.css", Type: "stylesheet"},
		{Item: "/public/dashboard.css", Type: "stylesheet"},
	}

	pg := controls.PageBody{}
	pg.Id = "mainpage"
	pg.Title = "Page Body"
	pg.HasThemeSwicther = true
	pg.Scripts = []string{
		"/public/bootstrap.bundle.min.js",
		"/public/dashboard.js",
		"/public/htmgo.js",
	}

	pg.GetBodyHeader = pg.SimpleNav
	pg.SidebarMenu = controls.SidebarMenu{
		MenuItems: []controls.NavMenuItem{
			controls.NewNavButton("show grid form", "/gridform", "#mainpage"),
			controls.NewNavButton("show card form", "/cardform", "#mainpage"),
			controls.NewNavHRule(),
		},
	}

	site.PageBody = pg

	//index
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderPage(request, writer, site.GetPage)
	})

	router.Get("/gridform", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, NewGF())
	})

	router.Get("/cardform", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, NewCF())
	})

	router.Get("/edititem/{id}", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderIPartial(request, writer, NewSillyForm())
	})

	http.ListenAndServe(":3000", router)
}

func NewCF() forms.CardGrid {
	gf := forms.CardGrid{}
	gf.CardRows = []controls.CardRow{
		NewCardRow(
			[]controls.Card{
				{
					Header: controls.CardHeader{
						Title: "Row 1 - Card1",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				{
					Header: controls.CardHeader{
						Title: "Row 1 - Card2",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				{
					Header: controls.CardHeader{
						Title: "Row 1 - Card3",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				/*
					{
						Header: controls.CardHeader{
							Title: "Row 1 - Card4",
						},
						Body: controls.CardBody{
							DrawControl: NewCardBody,
						},
						Footer: controls.Empty{},
					},
				*/
			},
		),
		NewCardRow(
			[]controls.Card{
				{
					Header: controls.CardHeader{
						Title: "Row 2 - Card1",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				{
					Header: controls.CardHeader{
						Title: "Row 2 - Card2",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				{
					Header: controls.CardHeader{
						Title: "Row 2 - Card3",
					},
					Body: controls.CardBody{
						DrawControl: NewCardBody,
					},
					Footer: controls.Empty{},
				},
				/*
					{
						Header: controls.CardHeader{
							Title: "Row 2 - Card4",
						},
						Body: controls.CardBody{
							DrawControl: NewCardBody,
						},
						Footer: controls.Empty{},
					},
				*/
			},
		),
	}
	return gf
}

func NewCardBody(Id string) *h.Element {
	return h.H1F("I am the card body")
}

func NewCardRow(cards []controls.Card) controls.CardRow {
	return controls.CardRow{
		CardsPerRow: controls.ThreePerRow,
		Cards:       cards,
	}
}

func GetBreadCrumbs() controls.Breadcrumbs {
	bc := controls.Breadcrumbs{}
	bc.Crumbs = []controls.Breadcrumb{
		{
			Text: "Home",
			Href: "/",
		},
		{
			Text: "edit",
			Href: "/",
		},
	}
	return bc
}

func NewGF() forms.GridForm {
	gf := forms.GridForm{}
	gf.Title = "A grid form"
	gf.HasUpdateTime = true
	gf.GetTable = getTable

	gf.Buttons = []controls.Button{
		{
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
					Name:  bootstrap.AttDataToggle,
					Value: "modal",
				},
				{
					Name:  bootstrap.AttDataTarget,
					Value: "#topmodal",
				},
			},
		},
	}

	//gf.BreadCrumbs = GetBreadCrumbs
	return gf
}

func NewSillyForm() forms.EditForm {
	sf := forms.EditForm{}
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
			Id:    "itemtexta",
			Title: "Type Something",
			Name:  "itemtexta",
			Cntrl: controls.TextArea{
				Rows:    20,
				Columns: 20,
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
		DrawControl: GetTableRows,
	}
	return tbl
}

func GetTableRows(Id string) *h.Element {
	rows := []string{}

	for i := 0; i < 100; i++ {
		rows = append(rows, fmt.Sprintf("row%d", i))
	}
	return h.List(rows, MakeCells)
}

func MakeCells(item string, index int) *h.Element {

	id := fmt.Sprintf("testid%d", index)

	cbox := controls.NewCheckedCheckbox(id+" displayname", id+"-chk", id+"-chk-name", "value1")

	if index%2 == 0 {
		cbox = controls.NewUnCheckedCheckbox(id+" displayname", id+"-chk", id+"-chk-name", "value2")
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
