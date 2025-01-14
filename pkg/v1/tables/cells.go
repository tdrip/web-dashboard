package tables

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

func GetActionButtonCell(editurl string, delurl string) *h.Element {
	if len(delurl) > 0 && len(editurl) > 0 {
		return h.Td(
			getEditBUtton(editurl),
			getDelBUtton(delurl),
		)
	}

	if len(delurl) == 0 && len(editurl) > 0 {
		return h.Td(
			getEditBUtton(editurl),
		)
	}
	if len(delurl) > 0 && len(editurl) == 0 {
		return h.Td(
			getDelBUtton(delurl),
		)
	}

	return h.Td(h.Text(" "))
}

func getEditBUtton(modalurl string) *h.Element {
	return h.Button(
		h.Class(bootstrap.Button, bootstrap.ButtonPrimary),
		h.Get(modalurl),
		h.Attribute(hx.TargetAttr, "#top-modal"),
		h.Attribute(hx.TriggerAttr, "click"),
		h.Attribute("data-bs-trigger", "modal"),
		h.Attribute("data-bs-target", "#top-modal"),
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
