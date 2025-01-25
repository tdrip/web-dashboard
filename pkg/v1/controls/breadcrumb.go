package controls

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/atts"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type Breadcrumb struct {
	BaseControl
	Attributes  []*h.AttributeR
	Classes     []string
	Text        string
	IsActive    bool
	Href        string
	DrawControl DrawControl
}

func (ctrl Breadcrumb) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Breadcrumb) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Breadcrumb) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Breadcrumb) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Breadcrumb) ToHTML() *h.Element {
	/*
	   	<li class="breadcrumb-item"><a href="#">Library</a></li>
	    <li class="breadcrumb-item active" aria-current="page">Data</li>
	*/
	if ctrl.IsActive {
		return h.Li(
			h.Class(bootstrap.BreadcrumbItem, "active"),
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			h.Text(ctrl.Text),
		)
	}

	if ctrl.DrawControl != nil {
		return h.Li(
			h.Class(bootstrap.BreadcrumbItem),
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			ctrl.DrawControl(),
		)
	}
	return h.Li(
		h.Class(bootstrap.BreadcrumbItem),
		h.Class(ctrl.Classes...),
		h.AttributeList(ctrl.Attributes...),
		h.A(
			h.Attribute(atts.Href, ctrl.Href),
			h.Text(ctrl.Text),
		),
	)
}

func ListBreadcrumb(ctrl Breadcrumb, index int) *h.Element {
	return ctrl.ToHTML()
}

func RenderBreadCrumbs(Crumbs []Breadcrumb) *h.Element {
	last := len(Crumbs) - 1
	fmt.Printf("last : %d", last)
	updated := []Breadcrumb{}
	for index, cr := range Crumbs {
		if index == last {
			fmt.Printf("IsActive : %d", index)
			cr.IsActive = true
		}
		updated = append(updated, cr)
	}
	return h.List(updated, ListBreadcrumb)
}
