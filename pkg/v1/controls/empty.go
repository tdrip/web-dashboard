package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Empty struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (ctrl Empty) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Empty) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Empty) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Empty) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Empty) ToHTML() *h.Element {
	return h.Empty()
}
