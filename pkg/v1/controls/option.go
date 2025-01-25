package controls

import (
	"strings"

	"github.com/maddalax/htmgo/framework/h"
)

type Option struct {
	BaseControl
	Text       string
	Value      string
	Selected   bool
	Attributes []*h.AttributeR
	Classes    []string
}

func NewSimpleOption(v string, selectedv string) Option {
	return Option{
		Text:     v,
		Value:    v,
		Selected: strings.EqualFold(v, selectedv),
	}
}

func (ctrl Option) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Option) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Option) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Option) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Option) ToHTML() *h.Element {
	atts := ctrl.Attributes

	atts = append(atts, &h.AttributeR{Name: "value", Value: ctrl.Value})

	if ctrl.Selected {
		atts = append(atts, &h.AttributeR{Name: "selected", Value: "selected"})
	}

	return h.Option(
		h.Class(ctrl.Classes...),
		h.AttributeList(atts...),
		h.Text(ctrl.Text),
	)
}

func ListOptions(ctrl Option, index int) *h.Element {
	return ctrl.ToHTML()
}
