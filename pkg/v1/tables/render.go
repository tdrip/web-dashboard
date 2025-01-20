package tables

import (
	"time"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"

	bootstrap "github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

const layout = "01-02-2006 15:04:05"

type TableRender interface {
	HasTitle() bool
	GetTitle() string
	HasNewButton() bool
	HasUpdateTime() bool
	GetModalCreateUrl() string
	GetModalCreateId() string
	GetTable() controls.Table
}

func RenderTable(tr TableRender) *h.Partial {
	tbl := tr.GetTable()
	return h.NewPartial(
		h.Div(
			checkHasTitle(tr),
			checkUpdateTime(tr),
			checkGetNew(tr),
			h.Div(
				h.Class(bootstrap.Row),
				h.Div(
					h.Class(bootstrap.Col),
					tbl.ToHTML(),
				),
			),
		),
	)
}

func checkUpdateTime(tr TableRender) *h.Element {
	if !tr.HasUpdateTime() {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col),
			h.Pf("Fetched: %s", time.Now().Format(layout)),
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

	newbtn := controls.Button{
		Text: "New",
		Classes: []string{
			bootstrap.Button,
			bootstrap.ButtonSuccess,
		},
		GetUrl: tr.GetModalCreateUrl(),
		Attributes: []*h.AttributeR{
			{
				Name:  hx.TargetAttr,
				Value: tr.GetModalCreateId(),
			},
			{
				Name:  "data-bs-toggle",
				Value: "modal",
			},
			{
				Name:  "data-bs-target",
				Value: tr.GetModalCreateId(),
			},
		},
	}

	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col, "d-grid", "gap-2", "d-md-flex", "justify-content-md-end"),
			newbtn.ToHTML(),
		),
	)
}
