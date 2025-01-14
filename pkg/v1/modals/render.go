package modals

import (
	"github.com/maddalax/htmgo/framework/h"
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
