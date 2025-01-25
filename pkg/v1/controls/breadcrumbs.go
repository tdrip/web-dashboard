package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/atts"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type Breadcrumbs struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Crumbs     []Breadcrumb
}

func (ctrl Breadcrumbs) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Breadcrumbs) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Breadcrumbs) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Breadcrumbs) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Breadcrumbs) ToHTML() *h.Element {
	return h.Nav(
		h.Attribute(atts.AriaLabel, bootstrap.Breadcrumb),
		h.Ol(
			h.Class(bootstrap.Breadcrumb),
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			RenderBreadCrumbs(ctrl.Crumbs),
		),
	)
}

func ListBreadcrumbs(ctrl Breadcrumbs, index int) *h.Element {
	return ctrl.ToHTML()
}
