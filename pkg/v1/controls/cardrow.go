package controls

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
)

type CardPerRow int

const (
	OnePerRow CardPerRow = iota
	TwoPerRow
	ThreePerRow
	FourPerRow
)

type CardRow struct {
	BaseControl
	CardsPerRow CardPerRow
	Cards       []Card
	Attributes  []*h.AttributeR
	Classes     []string
}

func (ctrl CardRow) GetClasses() []string {
	return ctrl.Classes
}

func (ctrl CardRow) SetClasses(classes []string) BaseControl {
	ctrl.Classes = SetClasses(ctrl, classes)
	return ctrl
}

func (ctrl CardRow) GetAtts() []*h.AttributeR {
	return ctrl.Attributes
}

func (ctrl CardRow) SetAtts(atts []*h.AttributeR) BaseControl {
	ctrl.Attributes = SetAtts(ctrl, atts)
	return ctrl
}

func (ctrl CardRow) ToHTML() *h.Element {
	return h.Div(
		h.Class(bootstrap.Row),
		h.AttributeList(ctrl.Attributes...),
		h.List(ctrl.Cards, ctrl.CardItems),
	)
}

func (ctrl CardRow) CardItems(card Card, index int) *h.Element {
	colclass := ""
	switch ctrl.CardsPerRow {
	case OnePerRow:
		colclass = "col-sm-12 mb-6 mb-sm-0"
	case TwoPerRow:
		colclass = "col-sm-6 mb-3 mb-sm-0"
	case ThreePerRow:
		colclass = "col-sm-4 mb-2 mb-sm-0"
	case FourPerRow:
		colclass = "col-sm-3 mb-2 mb-sm-0"
	}
	return h.Div(
		h.Class(colclass),
		card.ToHTML(),
	)
}

func ListCardRows(cr CardRow, index int) *h.Element {
	return cr.ToHTML()
}
