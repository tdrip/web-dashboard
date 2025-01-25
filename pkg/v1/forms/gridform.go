package forms

import (
	"fmt"
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
	Buttons       []controls.Button
	GetTable      GetTable
	HasUpdateTime bool
	IsForm        bool
	Attributes    []*h.AttributeR
	Classes       []string
}

func (ctrl GridForm) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ctrl.GetFormData == nil {
		return ctrl
	}
	return ctrl.GetFormData(ctx, ctrl)
}

func (ctrl GridForm) Render() *h.Partial {
	tbl := ctrl.GetTable()

	if ctrl.IsForm {
		return h.NewPartial(
			h.Div(
				h.Form(
					h.Class(ctrl.Classes...),
					h.AttributeList(ctrl.Attributes...),
					h.Div(
						h.Class("row", "d-flex", "justify-content-between", "flex-wrap", "flex-md-nowrap", "align-items-center", "pt-3", "pb-2", "mb-3", "border-bottom"),
						checkHasTitle(ctrl),
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
			),
		)
	}
	return h.NewPartial(
		h.Div(
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			h.Div(
				h.Class("row", "d-flex", "justify-content-between", "flex-wrap", "flex-md-nowrap", "align-items-center", "pt-3", "pb-2", "mb-3", "border-bottom"),
				checkHasTitle(ctrl),
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

func checkHasTitle(ctrl GridForm) *h.Element {
	if len(ctrl.Title) == 0 {
		return h.Empty()
	}

	if ctrl.HasUpdateTime {
		return h.Div(
			h.Class(bootstrap.Col),
			h.H1F(ctrl.Title),
			h.Span(
				h.Class("badge", "text-bg-secondary"),
				h.Text(fmt.Sprintf("Fetched: %s", time.Now().Format(layout))),
			),
		)
	}
	return h.Div(
		h.Class(bootstrap.Col),
		h.H1F(ctrl.Title),
	)

}

func checkGetNew(ctrl GridForm) *h.Element {
	if len(ctrl.Buttons) == 0 {
		return h.Empty()
	}

	if len(ctrl.Buttons) == 1 {
		return h.Div(
			h.Class(bootstrap.Col, "d-grid", "gap-2", "d-md-flex", "justify-content-md-end"),
			ctrl.Buttons[0].ToHTML(),
		)
	}

	return h.Div(
		h.Class(bootstrap.Col, "d-grid", "gap-2", "d-md-flex", "justify-content-md-end"),
		h.Div(
			h.Class("btn-toolbar", "mb-2", "md-md-0"),
			h.Div(
				h.Class("btn-group", "me-2"),
				h.List(ctrl.Buttons, controls.ListButtons),
			),
		),
	)
}
