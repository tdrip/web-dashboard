package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type GetTableRows func() *h.Element

type TableBody struct {
	BaseControl
	Attributes   []*h.AttributeR
	Classes      []string
	GetTableRows GetTableRows
}

func (ctrl TableBody) ToHTML() *h.Element {
	return h.TBody(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		ctrl.GetTableRows(),
	)
}

func (ctrl TableBody) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl TableBody) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl TableBody) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl TableBody) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
