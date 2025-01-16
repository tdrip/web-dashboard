package controls

import "github.com/maddalax/htmgo/framework/h"

type Option struct {
	DisplayName string
	Value       string
	Selected    bool
}

func (opt Option) ToHTML() *h.Element {

	if opt.Selected {
		return h.Option(
			h.Attribute("value", opt.Value),
			h.Attribute("selected", "selected"),
			h.Text(opt.DisplayName),
		)
	}

	return h.Option(
		h.Attribute("value", opt.Value),
		h.Text(opt.DisplayName),
	)
}
