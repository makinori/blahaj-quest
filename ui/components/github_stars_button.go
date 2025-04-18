package components

import (
	_ "embed"

	. "github.com/makinori/blahaj-quest/common"
	. "github.com/makinori/blahaj-quest/ui/icons"
	. "github.com/makinori/blahaj-quest/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func GitHubStarsButton(
	r *RenderContext, stars int, url string, children ...Node,
) Node {
	var starsStr string

	if stars < 0 {
		starsStr = "?"
	} else {
		starsStr = FormatNumber(stars)
	}

	return A(
		Href(url),
		Class(SCSS(r, `
			outline: solid 2px rgba(255,255,255,0.5);
			background: hsl(0, 0%, 95%);
			color: #000;
			border-radius: 6px;
			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: center;
			font-size: 12px;
			font-weight: 600;
			padding: 0 8px;
			gap: 8px;
			height: 24px;
		`)),
		GitHubIcon(
			Class(SCSS(r, `
				height: 12px;
			`)),
		),
		Text("Star"),
		Div(
			Class(SCSS(r, `
				background: white;
				height: 100%;
				padding-left: 8px;
				display: flex;
				align-items: center;
				justify-content: center;
				border-left: solid 1px hsl(0, 0%, 85%);
			`)),
			Text(starsStr),
		),
	)
}
