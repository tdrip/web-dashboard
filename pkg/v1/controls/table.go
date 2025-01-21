package controls

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Table struct {
	BaseControl
	Attributes   []*h.AttributeR
	Classes      []string
	TableHeaders TableHeaders
	TableBody    TableBody
}

func (ctrl Table) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Table) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Table) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Table) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Table) ToHTML() *h.Element {
	return h.Table(
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		ctrl.TableHeaders.ToHTML(),
		ctrl.TableBody.ToHTML(),
	)
}
