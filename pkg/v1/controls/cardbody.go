package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type CardBody struct {
	BaseControl
	Attributes  []*h.AttributeR
	Classes     []string
	DrawControl DrawControl
}

func (ctrl CardBody) ToHTML() *h.Element {
	return h.Div(
		h.Class(bootstrap.CardBody),
		h.AttributeList(ctrl.Attributes...),
		ctrl.DrawControl(),
	)
}

func (ctrl CardBody) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl CardBody) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl CardBody) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl CardBody) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}
