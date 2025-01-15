package forms

import (
	"strings"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

func GetInputControl(id string, title string, name string, value string, readonly bool) *h.Element {
	if readonly {
		return h.Div(
			h.Class(bootstrap.MB3, bootstrap.Row),
			h.LabelFor(id, title),
			h.TextInput(
				h.TextInput(
					h.Class(bootstrap.FormControl),
					h.Attribute("id", id),
					h.Attribute("name", name),
					h.Attribute("value", value),
					h.Attribute("readonly", "true"),
				),
			),
		)
	}
	return h.Div(
		h.Class(bootstrap.MB3, bootstrap.Row),
		h.LabelFor(id, title),
		h.TextInput(
			h.TextInput(
				h.Class(bootstrap.FormControl),
				h.Attribute("id", id),
				h.Attribute("name", name),
				h.Attribute("value", value),
			),
		),
	)
}

type Option struct {
	DisplayName string
	Value       string
	Selected    bool
}

func GetSelectControl(id string, title string, name string, values []string, selectedv string) *h.Element {
	opts := []Option{}
	for _, v := range values {
		opts = append(opts, Option{DisplayName: v, Value: v, Selected: strings.EqualFold(v, selectedv)})
	}
	return GetSelectControls(id, title, name, opts)
}

func GetSelectControls(id string, title string, name string, values []Option) *h.Element {
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

func optionItems(value Option, index int) *h.Element {

	if value.Selected {
		return h.Option(
			h.Attribute("value", value.Value),
			h.Attribute("selected", "selected"),
			h.Text(value.DisplayName),
		)
	}

	return h.Option(
		h.Attribute("value", value.Value),
		h.Text(value.DisplayName),
	)
}
