package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

type Button struct {
	GetUrl   string
	Id       string
	HXTarget string
	BSTarget string
	BSToggle string
	Classes  []string
	Text     string
}

func (btn Button) ToHTML() *h.Element {
	return h.Button(
		h.Class(btn.Classes...),
		h.Get(btn.GetUrl),
		h.Attribute(hx.TargetAttr, btn.HXTarget),
		h.Attribute("data-bs-toggle", "modal"),
		h.Attribute("data-bs-target", btn.BSTarget),
		h.Text(btn.Text),
	)
}

/*
h.Button(
	h.Class(bootstrap.Button, bootstrap.ButtonSecondary),
	h.Attribute("type", "button"),
	h.Attribute("data-bs-dismiss", "modal"),
	h.Text("Close"),
),
*/
