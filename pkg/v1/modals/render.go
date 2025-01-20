package modals

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type ModalRender interface {
	GetId() string
	GetTitle() string
	DataCreateURL() string
	GetModalBody() *h.Element
}

func RenderModal(mr ModalRender) *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Class(bootstrap.ModalDialog, bootstrap.ModalDialogCentered),
			h.Div(
				h.Class(bootstrap.ModalContent),
				h.Form(
					h.Attribute("hx-post", mr.DataCreateURL()),
					h.Div(
						h.Class(bootstrap.ModalHeader),
						h.H5(
							h.Class(bootstrap.ModalTitle),
							h.Text(mr.GetTitle()),
						),
					),
					mr.GetModalBody(),
					h.Div(
						h.Class(bootstrap.ModalFooter),
						h.Button(
							h.Class(bootstrap.Button, bootstrap.ButtonSecondary),
							h.Attribute("type", "button"),
							h.Attribute("data-bs-dismiss", "modal"),
							h.Text("Close"),
						),

						h.Button(
							h.Class(bootstrap.Button, bootstrap.ButtonSuccess),
							h.Attribute("data-bs-dismiss", "modal"),
							h.Text("Save"),
						),
					),
				),
			),
		),
	)
}
