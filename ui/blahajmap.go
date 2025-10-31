package ui

import (
	"context"
	_ "embed"
	"encoding/json"
	"log/slog"

	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohtml"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed blahajmap.js
	blahajMapJS string
)

func BlahajMap(ctx context.Context) Node {
	// TODO: error if more than one?

	// TODO: compress this better
	blahajData, err := json.Marshal(data.Blahaj.Current)
	if err != nil {
		slog.Error("failed to marshal blahaj data", "err", err)
	}

	return Div(
		Class(goemo.SCSS(ctx, `
			width: 100%;
			height: 100%;
			position: relative;

			#map {
				width: 100%;
				height: 100%;
			}
		`)),
		Div(ID("map")),
		emohtml.VStack(ctx,
			emohtml.StackSCSS(`
				position: absolute;
				z-index: 200;
				top: 8px;
				left: 8px;
				gap: 8px;
				user-select: none;
				align-items: flex-start;
			`),
			emohtml.HStack(ctx,
				emohtml.StackSCSS(`
					background: #fff;
					@include large-shadow;
					border-radius: 12px;
					font-size: 14px;
					gap: 2px;

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
			emohtml.VStack(ctx,
				emohtml.StackSCSS(`
					background: #fff;
					@include large-shadow;
					border-radius: 12px;
					font-size: 14px;
					gap: 0px;
					padding: 8px 0;

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
		// TODO: minify?
		Script(Raw(
			"(async ()=>{\n"+
				"const blahajData = "+string(blahajData)+";\n"+
				blahajMapJS+
				"\n})()")),
	)
}
