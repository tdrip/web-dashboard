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

func (ti TextInput) GetClasses() []string {
	return ti.Classes
}

func (ti TextInput) SetClassses(classes []string) BaseControl {
	ti.Classes = SetClassses(ti, classes)
	return ti
}

func (ti TextInput) GetAtts() []*h.AttributeR {
	return ti.Attributes
}

func (ti TextInput) SetAtts(atts []*h.AttributeR) BaseControl {
	ti.Attributes = SetAtts(ti, atts)
	return ti
}

func (ti TextInput) ToHTML() *h.Element {

	if ti.ReadOnly {
		ti.Attributes = SetAtt(ti, "readonly", "true")
	}

	return h.TextInput(
		h.Class(bootstrap.FormControl),
		h.AttributeList(ti.Attributes...),
	)

}
