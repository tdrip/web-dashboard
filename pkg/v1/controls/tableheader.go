package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type TableHeader struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (ctrl TableHeader) ToHTML() *h.Element {
	return h.Th(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		h.Text(ctrl.Text),
	)
}

func (ctrl TableHeader) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl TableHeader) SetClassses(classes []string) BaseControl {
	ctrl.Classes = SetClassses(ctrl, classes)
	return ctrl
}

func (ctrl TableHeader) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl TableHeader) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
