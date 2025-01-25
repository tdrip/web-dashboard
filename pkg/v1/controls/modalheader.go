package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type ModalHeader struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Title      string
}

func (ctrl ModalHeader) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl ModalHeader) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl ModalHeader) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl ModalHeader) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl ModalHeader) ToHTML() *h.Element {
	if len(ctrl.Title) == 0 {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.ModalHeader),
		h.H5(
			h.Class(bootstrap.ModalTitle),
			h.Text(ctrl.Title),
		),
	)
}
