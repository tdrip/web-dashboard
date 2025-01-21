package controls

import "github.com/maddalax/htmgo/framework/h"

func GetCheckBoxCell(item Checkbox) *h.Element {
	return h.Td(
		item.ToHTML(),
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

func GetButtonCell(buttons []Button) *h.Element {
	return h.Td(
		h.List(buttons, renderButtons),
	)
}

func renderButtons(item Button, index int) *h.Element {
	return item.ToHTML()
}
