package tables

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	"github.com/tdrip/web-dashboard/pkg/v1/controls"
)

func GetButtonCell(buttons []controls.Button) *h.Element {
	return h.Td(
		h.List(buttons, renderButtons),
	)
}

func renderButtons(item controls.Button, index int) *h.Element {
	return item.ToHTML()
}

func GetActionButtonCell(editurl string, delurl string, id string) *h.Element {
	if len(delurl) > 0 && len(editurl) > 0 {
		return h.Td(
			getEditBUtton(editurl, id),
			getDelBUtton(delurl),
		)
	}

	if len(delurl) == 0 && len(editurl) > 0 {
		return h.Td(
			getEditBUtton(editurl, id),
		)
	}
	if len(delurl) > 0 && len(editurl) == 0 {
		return h.Td(
			getDelBUtton(delurl),
		)
	}

	return h.Td(h.Text(" "))
}

func getEditBUtton(modalurl string, id string) *h.Element {
	return h.Button(
		h.Class(bootstrap.Button, bootstrap.ButtonPrimary),
		h.Get(modalurl),
		h.Attribute(hx.TargetAttr, id),
		h.Attribute(hx.TriggerAttr, "click"),
		h.Attribute("data-bs-trigger", "modal"),
		h.Attribute("data-bs-target", id),
		h.Text("Edit"),
	)
}

func getDelBUtton(delurl string) *h.Element {
	return h.Button(
		h.Class(bootstrap.Button, bootstrap.ButtonDanger),
		h.Attribute(hx.DeleteAttr, delurl),
		h.Text("Delete"),
	)
}

func GetTextCell(item string) *h.Element {
	return h.Td(
		h.Text(item),
	)
}

func GetListCell(items []string) *h.Element {
	return h.Td(
		h.List(items, getListItems),
	)
}

func getListItems(item string, index int) *h.Element {
	return h.Li(h.Text(item))
}
