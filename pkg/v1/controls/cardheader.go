package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type CardHeader struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Title      string
}

func (ctrl CardHeader) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl CardHeader) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl CardHeader) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl CardHeader) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl CardHeader) ToHTML() *h.Element {
	if len(ctrl.Title) == 0 {
		return h.Empty()
	}
	return h.Div(
		h.Class(bootstrap.CardHeader),
		h.H5(
			h.Class(bootstrap.CardTitle),
			h.Text(ctrl.Title),
		),
	)
}
