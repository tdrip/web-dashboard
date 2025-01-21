package forms

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
	"github.com/tdrip/web-dashboard/pkg/v1/render"
)

type GetModal func() controls.Modal

type ModalDialog struct {
	render.IPartial
	Attributes  []*h.AttributeR
	Classes     []string
	GetFormData GetFormData
	GetModal    GetModal
}

func (ctrl ModalDialog) DataFromContext(ctx *h.RequestContext) render.IPartial {
	if ctrl.GetFormData == nil {
		return ctrl
	}
	return ctrl.GetFormData(ctx, ctrl)
}

func (ctrl ModalDialog) Render() *h.Partial {
	mdl := ctrl.GetModal()
	return h.NewPartial(
		mdl.ToHTML(),
	)
}
