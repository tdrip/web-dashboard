package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TableHeaders struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
	Headers    []TableHeader
}

func GetSimpleTableHeaders(vals []string) TableHeaders {

	ths := TableHeaders{}
	headers := []TableHeader{}
	for _, v := range vals {
		headers = append(headers, TableHeader{Text: v})
	}
	ths.Headers = headers
	return ths
}

func (ctrl TableHeaders) ToHTML() *h.Element {

	return h.THead(
		h.Tr(
			h.Class(ctrl.Classes...),
			h.AttributeList(ctrl.Attributes...),
			h.Text(ctrl.Text),
			h.List(ctrl.Headers, ListHeaderitems),
		),
	)
}

func (ctrl TableHeaders) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl TableHeaders) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl TableHeaders) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl TableHeaders) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func ListHeaderitems(th TableHeader, index int) *h.Element {
	atts := th.GetAtts()
	atts = append(atts, &h.AttributeR{Name: "scope", Value: bootstrap.Col})
	updatedcntrl := th.SetAtts(atts)
	return updatedcntrl.ToHTML()
}
