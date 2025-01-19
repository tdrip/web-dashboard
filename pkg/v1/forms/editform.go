package forms

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type GetFormData func(c *h.RequestContext, ip render.IPartial) render.IPartial

type EditForm struct {
	render.IPartial
	Controls    []FormControl
	Attributes  []*h.AttributeR
	Classes     []string
	GetFormData GetFormData
	Title       string
}

func (ef EditForm) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ef.GetFormData == nil {
		return ef
	}
	return ef.GetFormData(ctx, ef)
}

func (ef EditForm) Render() *h.Partial {
	return h.NewPartial(
		h.Div(
			ef.checkEFHasTitle(),
			h.Div(
				h.Class(ef.Classes...),
				h.AttributeList(ef.Attributes...),
				h.List(ef.Controls, ListFormControls),
			),
		),
	)
}

func (ef EditForm) checkEFHasTitle() *h.Element {
	if len(ef.Title) == 0 {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.Row),
		h.Div(
			h.Class(bootstrap.Col),
			h.H2(
				h.Text(ef.Title),
			),
		),
	)
}
