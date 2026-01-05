package ui

import (
	"context"
	_ "embed"

	"git.hotmilk.space/maki/foxlib/foxcss"
	"git.hotmilk.space/maki/foxlib/foxhtml"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BlahajMap(ctx context.Context) Node {
	return Div(
		Class(foxcss.Class(ctx, `
			width: 100%;
			height: 100%;
			position: relative;

			#map {
				width: 100%;
				height: 100%;
			}
		`)),
		Div(ID("map")),
		foxhtml.VStack(ctx,
			foxhtml.StackCSS(`
				position: absolute;
				z-index: 200;
				top: 8px;
				left: 8px;
				gap: 8px;
				user-select: none;
				align-items: flex-start;
			`),
			foxhtml.HStack(ctx,
				foxhtml.StackCSS(`
					background: #fff;
					border-radius: 12px;
					font-size: 14px;
					gap: 2px;

					`+largeBoxShadowCSS+`

					label {
						padding: 8px 12px;
						display: flex;
						align-items: center;
						gap: 8px;
					}

					.divider {
						width: 2px;
						background: #000;
						opacity: 0.1;
					}
				`),
				Label(
					Input(Type("radio"), Name("map-style"), Value("openfreemap")),
					Text("OpenFreeMap"),
				),
				Div(Class("divider")),
				Label(
					Input(Type("radio"), Name("map-style"), Value("openstreetmap")),
					Text("OpenStreetMap"),
				),
			),
			foxhtml.VStack(ctx,
				foxhtml.StackCSS(`
					background: #fff;
					border-radius: 12px;
					font-size: 14px;
					gap: 0px;
					padding: 8px 0;
					
					`+largeBoxShadowCSS+`

					label {
						padding: 2px 12px;
						display: inline-flex;
						align-items: center;
						gap: 8px;
					}
				`),
				Label(
					Input(Type("checkbox"), Name("map-layer"), Value("blahaj")),
					Text("blåhaj"),
				),
				Label(
					Input(Type("checkbox"), Name("map-layer"), Value("blahaj-heatmap")),
					Text("heatmap"),
				),
			),
		),
	)
}
