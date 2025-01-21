package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type GetModalFooter func() *h.Element

type ModalFooter struct {
	BaseControl
	Attributes     []*h.AttributeR
	Classes        []string
	GetModalFooter GetModalFooter
}

func (ctrl ModalFooter) ToHTML() *h.Element {
	classes := SetClasses(ctrl, []string{bootstrap.ModalFooter})
	return h.Div(
		h.Class(classes...),
		h.AttributeList(ctrl.Attributes...),
		ctrl.GetModalFooter(),
	)
}

func (ctrl ModalFooter) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl ModalFooter) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl ModalFooter) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl ModalFooter) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
