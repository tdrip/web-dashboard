package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Button struct {
	BaseControl
	GetUrl     string
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (ctrl Button) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Button) SetClassses(classes []string) BaseControl {
	ctrl.Classes = SetClassses(ctrl, classes)
	return ctrl
}

func (ctrl Button) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Button) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Button) ToHTML() *h.Element {

	if len(ctrl.GetUrl) == 0 {
		return h.Button(
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			h.Text(ctrl.Text),
		)
	}

	return h.Button(
		h.Class(ctrl.Classes...),
		h.Get(ctrl.GetUrl),
		h.AttributeList(ctrl.Attributes...),
		h.Text(ctrl.Text),
	)
}
