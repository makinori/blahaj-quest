package components

import (
	_ "embed"
	"fmt"
	"time"

	. "github.com/makinori/blahaj-quest/blahaj"
	. "github.com/makinori/blahaj-quest/common"
	. "github.com/makinori/blahaj-quest/ui/icons"
	. "github.com/makinori/blahaj-quest/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func lastUpdated(lastUpdated time.Time) string {
	t := time.Since(lastUpdated)

	h := t.Hours()
	m := t.Minutes()

	if m < 60 {
		return fmt.Sprintf("%.0f minutes ago", m)
	}

	if h < 24 {
		return fmt.Sprintf("%.0f hours ago", h)
	}

	return fmt.Sprintf("%.0f days ago", h/24)
}

func PageHeader(r *RenderContext, children ...Node) []Node {
	blahaj := r.Context.Value(BlahajDataContextKey).(BlahajData)
	githubStars := r.Context.Value(GitHubStarsContextKey).(int)

	storeCount := len(blahaj.Data)

	var blahajCount int
	for _, store := range blahaj.Data {
		blahajCount += store.Quantity
	}

	return []Node{
		Div(
			Class(SCSS(r, `
				display: flex;
				flex-direction: row;
				align-items: center;
				gap: 16px;
				width: 100%;
				height: 64px;
				background-color: #3c8ea7;
				color: #fff;
			`)),
			// logo
			Img(
				Src("/img/full-flipped.png"),
				Class(SCSS(r, `
					width: 195px;
					margin-top: -32px;
					margin-left: -96px;
					display: block;
					object-fit: cover;
				`)),
			),
			// title
			Div(
				Class(SCSS(r, `
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
				`)),
				P(
					Class(SCSS(r, `
						font-size: 24px;
						letter-spacing: -0.02em;
						font-weight: 800;
					`)),
					Text("blåhaj quest"),
				),
				A(
					Class(SCSS(r, `
						font-size: 16px;
						font-weight: 700;
						opacity: 0.7;
					`)),
					Text(ConfigURL),
					Href(ConfigURL),
				),
			),
			// info
			Div(
				Class(SCSS(r, `
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					margin-left: 16px;
					font-weight: 500;
				`)),
				P(
					B(Text(Plural(storeCount, "store"))),
					Text(" with "),
					B(Text(FormatNumber(blahajCount)+" blåhaj")), // blahajar?
				),
				P(
					Text("last updated: "),
					B(Text(lastUpdated(blahaj.Updated))),
				),
			),
			Div(
				Class(SCSS(r, `flex-grow: 1`)),
			),
			// end of header
			Div(
				Class(SCSS(r, `
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					margin-right: 16px;
					gap: 6px;
					margin-top: -4px;
				`)),
				A(
					Href(ConfigGitHubURL),
					Class(SCSS(r, `
						display: flex;
						flex-direction: row;
						align-items: center;
						justify-content: center;
						font-size: 14px;
						font-weight: 600;
						gap: 6px;
					`)),
					Fa6Code(
						Class(SCSS(r, `
							fill: white;
							height: 12px;
						`)),
					),
					Text("see code"),
				),
				GitHubStarsButton(r, githubStars, ConfigGitHubURL),
			),
		),
		// shadow
		Div(
			Class(SCSS(r, `
				width: 100%;
				height: 8px;
				background-color: #307286;
			`)),
		),
	}
}
