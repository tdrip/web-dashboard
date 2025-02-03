package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type CardFooter struct {
	BaseControl
	Attributes  []*h.AttributeR
	Classes     []string
	DrawControl DrawControl
	Id          string
}

func (ctrl CardFooter) ToHTML() *h.Element {
	classes := SetClasses(ctrl, []string{bootstrap.CardFooter})
	return h.Div(
		h.Class(classes...),
		h.AttributeList(ctrl.Attributes...),
		ctrl.DrawControl(ctrl.Id),
	)
}

func (ctrl CardFooter) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl CardFooter) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl CardFooter) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl CardFooter) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
