package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Badge struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (ctrl Badge) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Badge) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Badge) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Badge) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Badge) ToHTML() *h.Element {
	return h.Span(
		h.Class("badge", "text-bg-secondary"),
		h.Text(ctrl.Text),
	)
}

func ListBadges(ctrl Badge, index int) *h.Element {
	return ctrl.ToHTML()
}
