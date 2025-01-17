package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type TextInput struct {
	Id       string
	ReadOnly bool
	Title    string
	Name     string
	Value    string
}

func (ti TextInput) ToHTML() *h.Element {
	if ti.ReadOnly {
		return h.TextInput(
			h.Class(bootstrap.FormControl),
			h.Attribute("id", ti.Id),
			h.Attribute("name", ti.Name),
			h.Attribute("value", ti.Value),
			h.Attribute("readonly", "true"),
		)
	}
	return h.TextInput(
		h.Class(bootstrap.FormControl),
		h.Attribute("id", ti.Id),
		h.Attribute("name", ti.Name),
		h.Attribute("value", ti.Value),
	)

}
