package ui

import (
	"context"
	"net/http"
	"strconv"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/blahaj-quest/ui/icons"
	"github.com/makinori/blahaj-quest/util"
	"github.com/makinori/foxlib/foxcss"
	"github.com/makinori/foxlib/foxhtml"
	"github.com/makinori/foxlib/foxhttp"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BlahajHeader(ctx context.Context, r *http.Request) Node {
	splitWidth := "770px"

	totalBlahaj := 0
	for i := range data.Blahaj.Current {
		totalBlahaj += data.Blahaj.Current[i].Quantity
	}

	infoEl := foxhtml.VStack(ctx,
		foxhtml.StackCSS(`
			gap: 0;
			font-weight: 500;
			b {
				font-weight: 700;
			}
		`),
		P(
			B(Text(util.FormatNumber(len(data.Blahaj.Current))+" stores")),
			Text(" with "),
			B(Text(util.FormatNumber(totalBlahaj)+" blåhaj")),
		),
		P(
			Text("last updated: "),
			// could use timediff but this function works
			B(Text(util.LastUpdated(data.Blahaj.Updated))),
		),
	)

	return Div(
		Class(foxcss.Class(ctx, `
			width: 100%;
			color: #fff;
			z-index: 100;
			`+largeBoxShadowCSS+`
		`)),
		foxhtml.HStack(ctx,
			foxhtml.StackCSS(`
				width: 100%;
				height: 64px;
				background-color: `+config.COLOR+`;
				align-items: center;
				gap: 0;
			`),
			Img(
				Src("/img/full-flipped.png"),
				Class(foxcss.Class(ctx, `
					width: 195px;
					margin-top: -32px;
					margin-left: -96px;
				`)),
			),
			foxhtml.VStack(ctx,
				foxhtml.StackCSS(`
					gap: 0;
					align-items: center;
					margin-left: 16px;
					h2 {
						font-weight: 800;
						font-size: 24px;
						letter-spacing: -0.02em;
					}
					a {
						font-weight: 700;
						opacity: 0.65;
					}
				`),
				H2(
					Text("blåhaj quest"),
				),
				A(
					Text("https://blahaj.quest"),
					Href("https://blahaj.quest"),
				),
			),
			Div(
				Class(foxcss.Class(ctx, `
					margin-left: 32px;
					@media (width < `+splitWidth+`) {
						& {
							display: none;
						}
					}
				`)),
				infoEl,
			),
			Img(
				Src("/notabot.gif?"+foxhttp.NotABotURLQuery(r)),
			),
			Div(Style("flex-grow: 1")),
			foxhtml.VStack(ctx,
				foxhtml.StackCSS(`
					align-items: center;
					gap: 0;
					margin-right: 24px;
					font-size: 14px;
				`),
				P(Text("made by")),
				foxhtml.HStack(ctx,
					foxhtml.StackCSS(`
						align-items: center;
						font-weight: 700;
						gap: 4px;
					`),
					A(Href("https://maki.cafe"), Text("maki")),
					Img(Src("/img/trans-heart.png"), Height("20")),
				),
			),
			A(
				Href(config.GITHUB_URL),
				foxhtml.HStack(ctx,
					foxhtml.StackCSS(`
						background-color: #fff;
						color: #000;
						outline: solid 2px rgba(255,255,255,0.5);
						border-radius: 6px;
						overflow: hidden;
						font-size: 12px;
						font-weight: 600;
						margin-right: 24px;
						align-items: center;
					`),
					foxhtml.HStack(ctx,
						foxhtml.StackCSS(`
							background-color: #edf2f7;
							padding: 5px 8px;
							border-right: solid 1px #cbd5e0;
							align-items: center;
						`),
						icons.GitHub(
							Height("1em"),
						),
						Text("Star"),
					),
					P(
						Class(foxcss.Class(ctx, `
							margin-right: 8px;
						`)),
						Text(strconv.Itoa(data.GitHubStars.Current)),
					),
				),
			),
		),
		foxhtml.HStack(ctx,
			foxhtml.StackCSS(`
				width: 100%;
				height: 64px;
				background-color: `+config.COLOR_LIGHTER+`;
				align-items: center;
				justify-content: center;
				@media (width >= `+splitWidth+`) {
					& {
						display: none;
					}
				}
			`),
			infoEl,
		),
		Div(Class(foxcss.Class(ctx, `
			width: 100%;
			height: 8px;
			background-color: `+config.COLOR_DARKER+`;
		`))),
	)

}
