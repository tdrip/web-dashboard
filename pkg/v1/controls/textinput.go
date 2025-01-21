package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TextInput struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	ReadOnly   bool
}

func (ctrl TextInput) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl TextInput) SetClassses(classes []string) BaseControl {
	ctrl.Classes = SetClassses(ctrl, classes)
	return ctrl
}

func (ctrl TextInput) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl TextInput) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl TextInput) ToHTML() *h.Element {

	if ctrl.ReadOnly {
		ctrl.Attributes = SetAtt(ctrl, "readonly", "true")
	}

	return h.TextInput(
		h.Class(bootstrap.FormControl),
		h.AttributeList(ctrl.Attributes...),
	)

}
