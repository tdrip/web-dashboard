package controls

import "github.com/maddalax/htmgo/framework/h"

type BaseControl interface {
	GetClasses() []string
	SetClassses(classes []string) BaseControl
	GetAtts() []*h.AttributeR
	SetAtts(atts []*h.AttributeR) BaseControl
	ToHTML() *h.Element
}

type BaseControlImp struct {
	BaseControl
	Attributes []*h.AttributeR
	Classes    []string
	Text       string
}

func (bci BaseControlImp) ToHTML() *h.Element {
	return h.Empty()
}

func (bci BaseControlImp) GetClasses() []string {
	return bci.Classes
}

func (bci BaseControlImp) SetClassses(classes []string) BaseControl {
	nclasses := bci.Classes
	nclasses = append(nclasses, classes...)
	bci.Classes = nclasses
	return bci
}

func (bci BaseControlImp) GetAtts() []*h.AttributeR {
	return bci.Attributes
}

func (bci BaseControlImp) SetAtts(atts []*h.AttributeR) BaseControl {
	natts := bci.Attributes
	natts = append(natts, atts...)
	bci.Attributes = natts
	return bci
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
