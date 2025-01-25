package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type Card struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Header     BaseControl
	Body       BaseControl
	Footer     BaseControl
}

func (ctrl Card) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Card) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Card) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Card) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Card) ToHTML() *h.Element {
	return h.Div(
		h.Class(bootstrap.Card),
		ctrl.Header.ToHTML(),
		ctrl.Body.ToHTML(),
		ctrl.Footer.ToHTML(),
	)
}
