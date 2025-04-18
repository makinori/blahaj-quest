package render

import (
	"context"

	. "maragu.dev/gomponents"
)

type RenderContext struct {
	Context context.Context
	SCSS    map[string]string
	HeadJS  map[string]string
	BodyJS  map[string]string
}

func Render(
	ctx context.Context,
	layout func(*RenderContext, ...Node) Node,
	page func(*RenderContext) Node,
) string {
	ensureSass()

	r := RenderContext{
		Context: ctx,
		SCSS:    map[string]string{},
		HeadJS:  map[string]string{},
		BodyJS:  map[string]string{},
	}

	html := Group{
		layout(&r, page(&r)),
	}.String()

	return html
}
