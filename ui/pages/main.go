package pages

import (
	. "github.com/makinori/blahaj-quest/ui/components"
	. "github.com/makinori/blahaj-quest/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MainPage(r *RenderContext) Node {
	children := []Node{
		Class(SCSS(r, `
			display: flex;
			flex-direction: column;
			width: 100vw;
			height: 100vh;
		`)),
	}

	children = append(children, PageHeader(r)...)
	children = append(children, BlahajMap(r))

	return Div(children...)
}
