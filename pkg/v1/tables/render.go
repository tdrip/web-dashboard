package tables

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"

	bootstrap "github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TableRender interface {
	GetHeaders() []string
	HasTitle() bool
	GetTitle() string
	HasNewButton() bool
	GetModalCreateUrl() string
	GetModalCreateId() string
	GetModalEditUrl(string) string
	GetTableBody() *h.Element
}

func RenderTable(tr TableRender) *h.Partial {
	return h.NewPartial(
		h.Div(
			checkHasTitle(tr),
			checkGetNew(tr),
			h.Div(
				h.Class(bootstrap.Row),
				h.Div(
					h.Class(bootstrap.Col, "table-responsive", "small"),
					h.Table(
						h.Class(bootstrap.TableClass, "table-striped", "table-sm", "delete-row-example"),
						getTableHeaders(tr.GetHeaders()),
						tr.GetTableBody(),
					),
				),
			),
		),
	)
}

func checkHasTitle(tr TableRender) *h.Element {
	if !tr.HasTitle() {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col),
			h.H2(
				h.Text(tr.GetTitle()),
			),
		),
	)
}

func checkGetNew(tr TableRender) *h.Element {
	if !tr.HasNewButton() {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col, "d-grid", "gap-2", "d-md-flex", "justify-content-md-end"),
			h.Button(
				h.Class(bootstrap.Button, bootstrap.ButtonSuccss),
				h.Get(tr.GetModalCreateUrl()),
				h.Attribute(hx.TargetAttr, tr.GetModalCreateId()),
				h.Attribute("data-bs-toggle", "modal"),
				h.Attribute("data-bs-target", tr.GetModalCreateId()),
				h.Text("New"),
			),
		),
	)
}

func getTableHeaders(headers []string) *h.Element {
	return h.THead(
		h.Tr(
			h.List(headers, headeritems),
		),
	)
}

func headeritems(item string, index int) *h.Element {
	return h.Th(
		h.Attribute("scope", bootstrap.Col),
		h.Text(item),
	)
}
