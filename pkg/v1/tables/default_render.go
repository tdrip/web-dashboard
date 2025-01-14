package tables

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

type Headers func() []string
type CreateUrl func() string
type TableRows func() *h.Element

type DefaultRender struct {
	TableRender
	DataUrl        string
	DataDelUrl     string
	Headers        Headers
	Title          string
	ModalCreateUrl string
	ModalEditUrl   string
	GetUrl         string
	Id             string
	TableRows      TableRows
}

func NewDefaultRender(dataurl string, headers Headers, title string, mcurl string, meurl string, delurl string, id string, tr TableRows) DefaultRender {
	return DefaultRender{
		DataUrl:        dataurl,
		DataDelUrl:     delurl,
		Headers:        headers,
		Title:          title,
		ModalCreateUrl: mcurl,
		ModalEditUrl:   meurl,
		Id:             id,
		TableRows:      tr,
	}
}

func (dr DefaultRender) GetHeaders() []string {
	return dr.Headers()
}

func (dr DefaultRender) GetTitle() string {
	return dr.Title
}

func (dr DefaultRender) GetTableBody() *h.Element {
	return h.TBody(
		h.Attribute("id", dr.Id),
		h.Attribute(hx.TargetAttr, "closest tr"),
		h.Attribute(hx.SwapAttr, hx.SwapTypeOuterHtml),
		dr.TableRows(),
	)
}

func (dr DefaultRender) GetModalCreateUrl() string {
	return dr.ModalCreateUrl
}

func (dr DefaultRender) GetModalEditUrl(id string) string {
	return fmt.Sprintf("%s/%s", dr.ModalEditUrl, id)
}

func (dr DefaultRender) GetDataDelUrl(id string) string {
	return fmt.Sprintf("%s/%s", dr.DataDelUrl, id)
}
