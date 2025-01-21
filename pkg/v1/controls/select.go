package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Select struct {
	BaseControl
	Options    []Option
	Attributes []*h.AttributeR
	Classes    []string
}

func (ctrl Select) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Select) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Select) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Select) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Select) ToHTML() *h.Element {
	return h.Select(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		h.List(ctrl.Options, OptionItems),
	)
}

func OptionItems(opt Option, index int) *h.Element {
	return opt.ToHTML()
}
