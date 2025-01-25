package controls

import "github.com/maddalax/htmgo/framework/h"

type BaseControl interface {
	GetClasses() []string
	SetClasses(classes []string) BaseControl
	GetAtts() []*h.AttributeR
	SetAtts(atts []*h.AttributeR) BaseControl
	ToHTML() *h.Element
}

func SetClasses(ctrl BaseControl, classes []string) []string {
	nclasses := ctrl.GetClasses()
	nclasses = append(nclasses, classes...)
	return nclasses
}

func SetAtts(ctrl BaseControl, atts []*h.AttributeR) []*h.AttributeR {
	natts := ctrl.GetAtts()
	natts = append(natts, atts...)
	return natts
}

func SetAtt(ctrl BaseControl, Name string, Value string) []*h.AttributeR {
	return SetAttR(ctrl, &h.AttributeR{Name: Name, Value: Value})
}

func SetAttR(ctrl BaseControl, val *h.AttributeR) []*h.AttributeR {
	natts := ctrl.GetAtts()
	natts = append(natts, val)
	return natts
}

func GetAttId(Id string) *h.AttributeR {
	return &h.AttributeR{Name: "id", Value: Id}
}

func SetAttId(ctrl BaseControl, Id string) []*h.AttributeR {
	return SetAttR(ctrl, GetAttId(Id))
}

func GetAttName(Name string) *h.AttributeR {
	return &h.AttributeR{Name: "name", Value: Name}
}

func SetAttName(ctrl BaseControl, Name string) []*h.AttributeR {
	return SetAttR(ctrl, GetAttName(Name))
}

func GetAttValue(Value string) *h.AttributeR {
	return &h.AttributeR{Name: "value", Value: Value}
}

func SetAttValue(ctrl BaseControl, Value string) []*h.AttributeR {
	return SetAttR(ctrl, GetAttValue(Value))
}

func BaseControlItems(ctrl BaseControl, index int) *h.Element {
	return ctrl.ToHTML()
}
