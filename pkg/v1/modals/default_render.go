package modals

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type DefaultRender struct {
	ModalRender
	Id        string
	Title     string
	CreateURL string
}

func (dr DefaultRender) GetId() string {
	return dr.Id
}

func (dr DefaultRender) GetTitle() string {
	return dr.Title
}

func (dr DefaultRender) DataCreateURL() string {
	return dr.CreateURL
}

func (dr DefaultRender) PartialRender(c *h.RequestContext) *h.Partial {
	return RenderModal(dr)
}

func (dr DefaultRender) GetModalBody() *h.Element {
	return h.Div(
		h.Attribute("id", dr.GetId()),
		h.Class(bootstrap.Modal, bootstrap.ModalBlur, "fade"),
		h.Attribute("style", "display:none"),
		h.AriaHidden(false),
		h.TabIndex(-11),
		h.Div(
			h.Class(bootstrap.ModalDialog, bootstrap.ModalLG, bootstrap.ModalDialogCentered),
			h.Attribute("role", "document"),
			h.Div(
				h.Class(bootstrap.ModalContent),
			),
		),
	)
}
