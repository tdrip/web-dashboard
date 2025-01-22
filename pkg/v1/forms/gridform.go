package forms

import (
	"time"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

const layout = "01-02-2006 15:04:05"

type GetTable func() controls.Table

type GridForm struct {
	render.IPartial
	GetFormData   GetFormData
	Title         string
	NewButton     *controls.Button
	GetTable      GetTable
	HasUpdateTime bool
}

func (ctrl GridForm) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ctrl.GetFormData == nil {
		return ctrl
	}
	return ctrl.GetFormData(ctx, ctrl)
}

func (ctrl GridForm) Render() *h.Partial {
	tbl := ctrl.GetTable()
	return h.NewPartial(
		h.Div(
			h.Div(
				h.Class("d-flex", "justify-content-between", "flex-wrap", "flex-md-nowrap", "align-items-center", "pt-3", "pb-2", "mb-3", "border-bottom"),
				checkHasTitle(ctrl),
				checkUpdateTime(ctrl),
				checkGetNew(ctrl),
			),
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

func checkUpdateTime(ctrl GridForm) *h.Element {
	if !ctrl.HasUpdateTime {
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

func checkHasTitle(ctrl GridForm) *h.Element {
	if len(ctrl.Title) == 0 {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col),
			h.H1F(ctrl.Title),
		),
	)
}

func checkGetNew(ctrl GridForm) *h.Element {
	if ctrl.NewButton == nil {
		return h.Empty()
	}

	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col, "d-grid", "gap-2", "d-md-flex", "justify-content-md-end"),
			ctrl.NewButton.ToHTML(),
		),
	)
}
