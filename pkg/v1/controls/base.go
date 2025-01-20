package controls

import "github.com/maddalax/htmgo/framework/h"

type BaseControl interface {
	GetClasses() []string
	SetClassses(classes []string) BaseControl
	GetAtts() []*h.AttributeR
	SetAtts(atts []*h.AttributeR) BaseControl
	ToHTML() *h.Element
}

func SetClassses(bci BaseControl, classes []string) []string {
	nclasses := bci.GetClasses()
	nclasses = append(nclasses, classes...)
	return nclasses
}

func SetAtts(bci BaseControl, atts []*h.AttributeR) []*h.AttributeR {
	natts := bci.GetAtts()
	natts = append(natts, atts...)
	return natts
}
