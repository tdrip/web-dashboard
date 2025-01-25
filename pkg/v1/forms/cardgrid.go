package forms

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type CardGrid struct {
	render.IPartial
	GetFormData GetFormData
	CardRows    []controls.CardRow
	Attributes  []*h.AttributeR
	Classes     []string
}

func (ctrl CardGrid) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ctrl.GetFormData == nil {
		return ctrl
	}
	return ctrl.GetFormData(ctx, ctrl)
}

func (ctrl CardGrid) Render() *h.Partial {
	return h.NewPartial(
		h.List(ctrl.CardRows, controls.ListCardRows),
	)
}
