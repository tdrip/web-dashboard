package forms

import (
	"strings"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

func GetInputControl(id string, title string, name string, value string, readonly bool) *h.Element {
	if readonly {
		return h.Div(
			h.Class(bootstrap.MB3, bootstrap.Row),
			h.LabelFor(id, title),
			h.TextInput(

				h.Class(bootstrap.FormControl),
				h.Attribute("id", id),
				h.Attribute("name", name),
				h.Attribute("value", value),
				h.Attribute("readonly", "true"),
			),
		)
	}
	return h.Div(
		h.Class(bootstrap.MB3, bootstrap.Row),
		h.LabelFor(id, title),
		h.TextInput(
			h.Class(bootstrap.FormControl),
			h.Attribute("id", id),
			h.Attribute("name", name),
			h.Attribute("value", value),
		),
	)
}

func GetSelectControl(id string, title string, name string, values []string, selectedv string) *h.Element {
	opts := []controls.Option{}
	for _, v := range values {
		opts = append(opts, controls.Option{DisplayName: v, Value: v, Selected: strings.EqualFold(v, selectedv)})
	}
	return GetSelectControls(id, title, name, opts)
}

func GetSelectControls(id string, title string, name string, values []controls.Option) *h.Element {
	return h.Div(
		h.Class(bootstrap.MB3, bootstrap.Row),
		h.LabelFor(id, title),
		h.Select(
			h.Class(bootstrap.FormControl),
			h.Attribute("id", id),
			h.Attribute("name", name),
			h.List(values, optionItems),
		),
	)
}

func optionItems(value controls.Option, index int) *h.Element {
	return value.ToHTML()
}
