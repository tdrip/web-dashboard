package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	atts "github.com/tdrip/web-dashboard/pkg/v1/atts"
)

type Checkbox struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (ctrl Checkbox) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Checkbox) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Checkbox) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Checkbox) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Checkbox) ToHTML() *h.Element {
	id := ""
	for _, att := range ctrl.Attributes {
		if att.Name == atts.Id {
			id = att.Value
			break
		}
	}

	return h.Div(
		h.Class("form-check"),
		h.Checkbox(
			h.Class("form-check-input"),
			h.AttributeList(ctrl.Attributes...),
		),
		h.Label(
			h.Class("form-check-label"),
			h.Attribute(atts.For, id),
			h.Text(ctrl.Text),
		),
	)

}

func NewCheckedCheckbox(text string, id string, formname string, formvalue string) Checkbox {
	cbox := Checkbox{
		Text: text,
		Attributes: []*h.AttributeR{
			{
				Name:  atts.Id,
				Value: id,
			},
			{
				Name:  atts.Checked,
				Value: "",
			},
			{
				Name:  atts.Name,
				Value: formname,
			},
			{
				Name:  atts.Value,
				Value: formvalue,
			},
		},
	}
	return cbox
}

func NewUnCheckedCheckbox(text string, id string, formname string, formvalue string) Checkbox {
	cbox := Checkbox{
		Text: text,
		Attributes: []*h.AttributeR{
			{
				Name:  atts.Id,
				Value: id,
			},
			{
				Name:  atts.Name,
				Value: formname,
			},
			{
				Name:  atts.Value,
				Value: formvalue,
			},
		},
	}
	return cbox
}

func ListCheckboxes(ctrl Checkbox, index int) *h.Element {
	return ctrl.ToHTML()
}
