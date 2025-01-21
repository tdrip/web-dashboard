package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type GetModalBody func() *h.Element

type ModalBody struct {
	BaseControl
	Attributes   []*h.AttributeR
	Classes      []string
	GetModalBody GetModalBody
}

func (ctrl ModalBody) ToHTML() *h.Element {
	return h.TBody(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		ctrl.GetModalBody(),
	)
}

func (ctrl ModalBody) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl ModalBody) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl ModalBody) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl ModalBody) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
