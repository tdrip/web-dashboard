package controls

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TextArea struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	ReadOnly   bool
	Rows       int
	Columns    int
}

func (ctrl TextArea) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl TextArea) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl TextArea) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl TextArea) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl TextArea) ToHTML() *h.Element {

	if ctrl.ReadOnly {
		ctrl.Attributes = SetAtt(ctrl, "readonly", "true")
	}

	if ctrl.Rows < 0 {
		ctrl.Attributes = SetAtt(ctrl, "rows", fmt.Sprintf("%d", ctrl.Rows))
	}

	if ctrl.Columns < 0 {
		ctrl.Attributes = SetAtt(ctrl, "cols", fmt.Sprintf("%d", ctrl.Columns))
	}

	return h.TextArea(
		h.Class(bootstrap.FormControl),
		h.AttributeList(ctrl.Attributes...),
	)

}
