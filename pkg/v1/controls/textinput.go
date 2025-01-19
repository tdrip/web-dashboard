package controls

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TextInput struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
	Id         string
	ReadOnly   bool
	Name       string
	Value      string
}

func (ti TextInput) GetClasses() []string {
	return ti.Classes
}

func (ti TextInput) SetClassses(classes []string) BaseControl {
	//nclasses := ti.Classes
	//nclasses = append(nclasses, classes...)
	//ti.Classes = nclasses
	ti.Classes = SetClassses(ti, classes)
	return ti
}

func (ti TextInput) GetAtts() []*h.AttributeR {
	return ti.Attributes
}

func (ti TextInput) SetAtts(atts []*h.AttributeR) BaseControl {
	//natts := ti.Attributes
	//natts = append(natts, atts...)
	ti.Attributes = SetAtts(ti, atts)
	return ti
}

func (ti TextInput) ToHTML() *h.Element {

	fmt.Println("TOHTML Called")
	atts := ti.Attributes

	atts = append(atts, &h.AttributeR{Name: "id", Value: ti.Id})
	atts = append(atts, &h.AttributeR{Name: "name", Value: ti.Name})
	atts = append(atts, &h.AttributeR{Name: "value", Value: ti.Value})

	if ti.ReadOnly {
		atts = append(atts, &h.AttributeR{Name: "readonly", Value: "true"})
	}

	return h.TextInput(
		h.Class(bootstrap.FormControl),
		h.AttributeList(atts...),
	)

}
