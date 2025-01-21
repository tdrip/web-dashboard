package forms

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type GetFormData func(c *h.RequestContext, ip render.IPartial) render.IPartial

type EditForm struct {
	render.IPartial
	Controls    []FormControl
	Buttons     []controls.Button
	Attributes  []*h.AttributeR
	Classes     []string
	GetFormData GetFormData
	Title       string
}

func (ctrl EditForm) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ctrl.GetFormData == nil {
		return ctrl
	}
	return ctrl.GetFormData(ctx, ctrl)
}

func (ctrl EditForm) Render() *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Form(
				h.Class(ctrl.Classes...),
				h.AttributeList(ctrl.Attributes...),
				ctrl.checkHasTitle(),
				h.Div(
					h.List(ctrl.Controls, ListFormControls),
				),
				h.Div(
					h.Class(bootstrap.Row),
					h.Div(
						h.Class("d-flex", "gap-2", "justify-content-center", "py-5"),
						h.List(ctrl.Buttons, ListFormButtons),
					),
				),
			),
		),
	)
}

func (ctrl EditForm) checkHasTitle() *h.Element {
	if len(ctrl.Title) == 0 {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col),
			h.H2(
				h.Text(ctrl.Title),
			),
		),
	)
}
