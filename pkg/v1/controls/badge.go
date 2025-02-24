package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type BadgeColoring int

const (
	None BadgeColoring = iota
	Primary
	Secondary
	Success
	Danger
	Warning
	Info
	Light
	Dark
)

type Badge struct {
	BaseControl
	Attributes    []*h.AttributeR
	Classes       []string
	Text          string
	IsRoundedPill bool
	Coloring      BadgeColoring
}

func (ctrl Badge) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl Badge) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl Badge) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl Badge) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl Badge) ToHTML() *h.Element {

	colouring := ""
	switch ctrl.Coloring {
	case None:
		colouring = ""
	case Primary:
		colouring = bootstrap.TextBGPrimary
	case Secondary:
		colouring = bootstrap.TextBGSecondary
	case Success:
		colouring = bootstrap.TextBGSuccess
	case Danger:
		colouring = bootstrap.TextBGDanger
	case Warning:
		colouring = bootstrap.TextBGWarning
	case Info:
		colouring = bootstrap.TextBGInfo
	case Light:
		colouring = bootstrap.TextBGLight
	case Dark:
		colouring = bootstrap.TextBGDark
	}

	if ctrl.IsRoundedPill {
		return h.Span(
			h.Class(bootstrap.Badge, bootstrap.RoundedPill, colouring),
			h.Text(ctrl.Text),
		)
	}
	return h.Span(
		h.Class(bootstrap.Badge, colouring),
		h.Text(ctrl.Text),
	)
}

func ListBadges(ctrl Badge, index int) *h.Element {
	return ctrl.ToHTML()
}
