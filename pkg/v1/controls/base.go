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

func SetAtt(bci BaseControl, Name string, Value string) []*h.AttributeR {
	return SetAttR(bci, &h.AttributeR{Name: Name, Value: Value})
}

func SetAttR(bci BaseControl, val *h.AttributeR) []*h.AttributeR {
	natts := bci.GetAtts()
	natts = append(natts, val)
	return natts
}

func GetAttId(Id string) *h.AttributeR {
	return &h.AttributeR{Name: "id", Value: Id}
}

func SetAttId(bci BaseControl, Id string) []*h.AttributeR {
	return SetAttR(bci, GetAttId(Id))
}

func GetAttName(Name string) *h.AttributeR {
	return &h.AttributeR{Name: "name", Value: Name}
}

func SetAttName(bci BaseControl, Name string) []*h.AttributeR {
	return SetAttR(bci, GetAttName(Name))
}

func GetAttValue(Name string) *h.AttributeR {
	return &h.AttributeR{Name: "name", Value: Name}
}
func SetAttValue(bci BaseControl, Value string) []*h.AttributeR {
	return SetAttR(bci, GetAttValue(Value))
}
