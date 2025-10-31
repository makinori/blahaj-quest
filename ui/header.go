package ui

import (
	"context"
	"strconv"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/blahaj-quest/ui/icons"
	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohtml"
	"github.com/mergestat/timediff"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BlahajHeader(ctx context.Context) Node {
	splitWidth := "770px"

	infoEl := emohtml.VStack(ctx,
		emohtml.StackSCSS(`
			gap: 0;
			font-weight: 500;
			b {
				font-weight: 700;
			}
		`),
		P(
			B(Text("430 stores")),
			Text(" with "),
			B(Text("24,582 blåhaj")),
		),
		P(
			Text("last updated: "),
			B(Text(
				timediff.TimeDiff(data.Blahaj.Updated),
			)),
		),
	)

	return Div(
		Class(goemo.SCSS(ctx, `
			width: 100%;
			color: #fff;
		`)),
		emohtml.HStack(ctx,
			emohtml.StackSCSS(`
				width: 100%;
				height: 64px;
				background-color: `+config.COLOR+`;
				align-items: center;
				gap: 0;
			`),
			Img(
				Src("/img/full-flipped.png"),
				Class(goemo.SCSS(ctx, `
					width: 195px;
					margin-top: -32px;
					margin-left: -96px;
				`)),
			),
			emohtml.VStack(ctx,
				emohtml.StackSCSS(`
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
				Class(goemo.SCSS(ctx, `
					margin-left: 32px;
					@media (width < `+splitWidth+`) {
						display: none;
					}
				`)),
				infoEl,
			),
			Div(Style("flex-grow:1")),
			emohtml.VStack(ctx,
				emohtml.StackSCSS(`
					align-items: center;
					gap: 0;
					margin-right: 24px;
					font-size: 14px;
				`),
				P(Text("made by")),
				emohtml.HStack(ctx,
					emohtml.StackSCSS(`
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
				emohtml.HStack(ctx,
					emohtml.StackSCSS(`
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
					emohtml.HStack(ctx,
						emohtml.StackSCSS(`
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
						Class(goemo.SCSS(ctx, `
							margin-right: 8px;
						`)),
						Text(strconv.Itoa(data.GitHubStars.Current)),
					),
				),
			),
		),
		emohtml.HStack(ctx,
			emohtml.StackSCSS(`
				width: 100%;
				height: 64px;
				background-color: `+config.COLOR_LIGHTER+`;
				align-items: center;
				justify-content: center;
				@media (width >= `+splitWidth+`) {
					display: none;
				}
			`),
			infoEl,
		),
		Div(Class(goemo.SCSS(ctx, `
			width: 100%;
			height: 8px;
			background-color: `+config.COLOR_DARKER+`;
		`))),
	)

}
