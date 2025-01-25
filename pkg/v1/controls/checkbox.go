package controls

import (
	"github.com/maddalax/htmgo/framework/h"
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
		if att.Name == "id" {
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
			h.Attribute("for", id),
			h.Text(ctrl.Text),
		),
	)

}

func NewCheckedCheckbox(text string, id string, formname string, formvalue string) Checkbox {
	cbox := Checkbox{
		Text: text,
		Attributes: []*h.AttributeR{
			{
				Name:  "id",
				Value: id,
			},
			{
				Name:  "checked",
				Value: "",
			},
			{
				Name:  "name",
				Value: formname,
			},
			{
				Name:  "value",
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
				Name:  "id",
				Value: id,
			},
			{
				Name:  "name",
				Value: formname,
			},
			{
				Name:  "value",
				Value: formvalue,
			},
		},
	}
	return cbox
}

/*
<div class="form-check">
  <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault">
  <label class="form-check-label" for="flexCheckDefault">
    Default checkbox
  </label>
</div>
*/
