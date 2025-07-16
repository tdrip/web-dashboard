package render

import "github.com/maddalax/htmgo/framework/h"

type IPartial interface {
	DataFromContext(ctx *h.RequestContext) IPartial

	// Render Partial
	Render() *h.Partial
}
