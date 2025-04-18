package components

import (
	_ "embed"

	. "github.com/makinori/blahaj-quest/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BlahajMap(r *RenderContext) Node {
	mapStyles := []struct {
		Key   string
		Title string
	}{
		{Key: "maptiler", Title: "MapTiler"},
		{Key: "osm", Title: "OpenStreetMap"},
	}

	var mapStyleNodes []Node

	for i, mapStyle := range mapStyles {
		mapStyleNodes = append(mapStyleNodes,
			Input(
				Type("radio"), ID("map-style-"+mapStyle.Key),
				Name("map-style"), Value(mapStyle.Key),
			),
			Label(For("map-style-"+mapStyle.Key), Text(mapStyle.Title)),
		)

		if i < len(mapStyles)-1 {
			mapStyleNodes = append(mapStyleNodes,
				Div(Class(SCSS(r, `
					width: 2px;
					height: 100%;
					background-color: hsl(0, 0, 90%);
					margin: 0 3px;
				`))),
			)
		}
	}

	return Div(
		Class(SCSS(r, `
			background-color: #86ccfa;
			width: 100%;
			height: 100%;
			position: relative;
		`)),
		// map settings
		Div(
			Class(SCSS(r, `
				position: absolute;
				top: 8px;
				left: 8px;
			`)),
			// map style selector
			Div(append([]Node{
				Class(SCSS(r, `
					display: flex;
					flex-direction: row;
					align-items: center;
					justify-content: center;
					padding: 0 12px;
					height: 32px;
					background: white;
					border-radius: 12px;
					gap: 6px;
					font-size: 14px;

					input, label {
						cursor: pointer;
					}
				`))},
				mapStyleNodes...,
			)...),
		),
	)
}
