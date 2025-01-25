package forms

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

type GetBreadCrumbs func() controls.Breadcrumbs

func GetInputControl(id string, title string, name string, value string, readonly bool) *h.Element {

	fc := FormControl{}
	fc.Id = id
	fc.Title = title
	fc.Name = name

	ti := controls.TextInput{}
	ti.ReadOnly = readonly
	ti.Attributes = controls.SetAttValue(ti, value)
	fc.Cntrl = ti

	return ListFormControls(fc, -1)
}

func GetSelectControl(id string, title string, name string, values []string, selectedv string) *h.Element {
	opts := []controls.Option{}
	for _, v := range values {
		opts = append(opts, controls.NewSimpleOption(v, selectedv))
	}
	return GetSelectControls(id, title, name, opts)
}

func GetSelectControls(id string, title string, name string, values []controls.Option) *h.Element {
	fc := FormControl{}
	fc.Id = id
	fc.Title = title
	fc.Name = name

	sel := controls.Select{}
	sel.Options = values
	fc.Cntrl = sel
	return ListFormControls(fc, -1)
}
