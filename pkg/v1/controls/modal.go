package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type Modal struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Header     ModalHeader
	Body       ModalBody
	Footer     ModalFooter
}

func (ctrl Modal) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Modal) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Modal) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Modal) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Modal) ToHTML() *h.Element {
	return h.Div(
		h.Class(bootstrap.ModalDialog, bootstrap.ModalDialogCentered),
		h.Div(
			h.Class(bootstrap.ModalContent),
			h.Form(
				h.Class(ctrl.Classes...),
				h.AttributeList(ctrl.Attributes...),
				ctrl.Header.ToHTML(),
				ctrl.Body.ToHTML(),
				ctrl.Footer.ToHTML(),
			),
		),
	)
}

func GetEmptyModal(ModalID string) *h.Element {
	return h.Div(
		h.Attribute("id", ModalID),
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
