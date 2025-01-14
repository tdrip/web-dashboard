package modals

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type ModalRender interface {
	GetTitle() string
	DataCreateURL() string
	GetModalBody() *h.Element
}

func EmptyModal() *h.Element {
	return h.Div(
		h.Attribute("id", "top-modal"),
		h.Class("modal", "modal-blur", "fade"),
		h.Attribute("style", "display:none"),
		h.AriaHidden(false),
		h.TabIndex(-11),
		h.Div(
			h.Class("modal-dialog", "modal-lg", "modal-dialog-centered"),
			h.Attribute("role", "document"),
			h.Div(
				h.Class("modal-content"),
			),
		),
	)
}

func RenderModal(mr ModalRender) *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Class("modal-dialog", "modal-dialog-centered"),
			h.Div(
				h.Class("modal-content"),
				h.Form(
					h.Attribute("hx-post", mr.DataCreateURL()),
					h.Div(
						h.Class("modal-header"),
						h.H5(
							h.Class("modal-title"),
							h.Text(mr.GetTitle()),
						),
					),
					mr.GetModalBody(),
					h.Div(
						h.Class("modal-footer"),
						h.Button(
							h.Class(bootstrap.Button, bootstrap.ButtonSecondary),
							h.Attribute("type", "button"),
							h.Attribute("data-bs-dismiss", "modal"),
							h.Text("Close"),
						),

						h.Button(
							h.Class(bootstrap.Button, bootstrap.ButtonSuccss),
							h.Attribute("data-bs-dismiss", "modal"),
							h.Text("Save"),
						),
					),
				),
			),
		),
	)
}
