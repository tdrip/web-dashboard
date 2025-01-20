package forms

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

type FormControl struct {
	Id    string
	Name  string
	Title string
	Cntrl controls.BaseControl
}

func ListFormControls(fc FormControl, index int) *h.Element {
	updatedcntrl := fc.Cntrl

	// this form control has an ID and a Name
	atts := updatedcntrl.GetAtts()
	atts = append(atts, &h.AttributeR{Name: "id", Value: fc.Id})
	atts = append(atts, &h.AttributeR{Name: "name", Value: fc.Name})
	updatedcntrl = updatedcntrl.SetAtts(atts)

	// make the under lying control look good
	classes := updatedcntrl.GetClasses()
	classes = append(classes, bootstrap.FormControl)
	updatedcntrl = updatedcntrl.SetClassses(classes)
	fc.Cntrl = updatedcntrl

	fmt.Println("working on " + fc.Id)
	fmt.Printf("fc %v\n", fc)
	fmt.Printf("Cntrl %v\n", fc.Cntrl)
	fmt.Printf("updatedcntrl %v\n", updatedcntrl)

	return h.Div(
		h.Class(bootstrap.MB3, bootstrap.Row),
		h.LabelFor(fc.Id, fc.Title),
		updatedcntrl.ToHTML(),
	)
}

func ListFormButtons(btn controls.Button, index int) *h.Element {
	//
	classes := btn.GetClasses()
	classes = append(classes, "d-inline-flex")
	classes = append(classes, "align-items-center")
	updatedcntrl := btn.SetClassses(classes)
	return updatedcntrl.ToHTML()
}
