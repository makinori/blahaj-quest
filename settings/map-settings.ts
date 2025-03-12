import layers from "protomaps-themes-base";
import { create } from "zustand";

export enum MapStyle {
	MapTiler = "MapTiler",
	OpenStreetMap = "OpenStreetMap",
	// Protomaps = "Protomaps",
	// StamenWatercolor = "StamenWatercolor",
}

interface MapSettings {
	style: MapStyle;
	blahajLayer: boolean;
	heatmapLayer: boolean;
	setStyle: (style: MapStyle) => any;
	setBlahajLayer: (blahajLayer: boolean) => any;
	setHeatmapLayer: (heatmapLayer: boolean) => any;
}

export const useMapSettings = create<MapSettings>()(
	// persist(
	set => ({
		style: MapStyle.MapTiler,
		blahajLayer: true,
		heatmapLayer: true,
		setStyle: (style: MapStyle) => set({ style }),
		setBlahajLayer: (blahajLayer: boolean) => set({ blahajLayer }),
		setHeatmapLayer: (heatmapLayer: boolean) => set({ heatmapLayer }),
	}),
	// 	{
	// 		name: "blahaj-quest", // unique name
	// 		storage: createJSONStorage(() => localStorage),
	// 	},
	// ),
);

// light, dark, white, black, grayscale or debug
const pmLayers = layers("protomaps", "light");
// water color is too bright
const pmWaterLayer = pmLayers.find(layer => layer.id == "water");
if (pmWaterLayer != null) {
	pmWaterLayer.paint["fill-color"] = "#aad3df";
}

export const MapStyleConfigs: { [key in MapStyle]: any } = {
	// https://github.com/element-hq/element-web/blob/develop/config.sample.json
	MapTiler:
		"https://api.maptiler.com/maps/streets/style.json?key=fU3vlMsMn4Jb6dnEIFsx",
	OpenStreetMap: {
		version: 8,
		name: "OpenStreetMap Mapnik raster tiles (Default)",
		metadata: {
			"mapbox:autocomposite": true,
		},
		glyphs: "https://cdn.jsdelivr.net/gh/lukasmartinelli-alt/glfonts@gh-pages/fonts/{fontstack}/{range}.pbf",
		sources: {
			"osm-mapnik": {
				type: "raster",
				tiles: ["https://a.tile.openstreetmap.org/{z}/{x}/{y}.png"],
				tileSize: 256,
				attribution:
					"Basemap data <a href='https://www.osm.org' target=_blank>© OpenStreetMap contributors</a>",
			},
		},
		layers: [
			{
				id: "background",
				type: "background",
				paint: {
					"background-color": "rgba(0,0,0,0)",
				},
			},
			{
				id: "osm-mapnik",
				type: "raster",
				source: "osm-mapnik",
			},
		],
		owner: "OpenStreetMap India",
	},
	// Protomaps: {
	// 	version: 8,
	// 	glyphs: "https://cdn.protomaps.com/fonts/pbf/{fontstack}/{range}.pbf", // should self host this too
	// 	sources: {
	// 		protomaps: {
	// 			type: "vector",
	// 			// url: "pmtiles://https://build.protomaps.com/20231025.pmtiles",
	// 			tiles: [
	// 				"https://pmtiles.hotmilk.space/20231025/{z}/{x}/{y}.mvt",
	// 			],
	// 			maxzoom: 15,
	// 			attribution:
	// 				'<a href="https://protomaps.com">Protomaps</a> © <a href="https://openstreetmap.org">OpenStreetMap</a>',
	// 		},
	// 	},
	// 	layers: pmLayers,
	// },
	// StamenWatercolor: {
	// 	version: 8,
	// 	name: "OpenStreetMap Mapnik raster tiles (Default)",
	// 	metadata: {
	// 		"mapbox:autocomposite": true,
	// 	},
	// 	glyphs: "https://cdn.jsdelivr.net/gh/lukasmartinelli-alt/glfonts@gh-pages/fonts/{fontstack}/{range}.pbf",
	// 	sources: {
	// 		"osm-mapnik": {
	// 			type: "raster",
	// 			tiles: [
	// 				"https://stamen-tiles-c.a.ssl.fastly.net/watercolor/{z}/{x}/{y}.jpg",
	// 			],
	// 			tileSize: 256,
	// 			attribution:
	// 				"<a href='http://maps.stamen.com' target=_blank>© Stamen Maps</a> <a href='https://www.osm.org' target=_blank>© OpenStreetMap contributors</a>",
	// 		},
	// 	},
	// 	layers: [
	// 		{
	// 			id: "background",
	// 			type: "background",
	// 			paint: {
	// 				"background-color": "rgba(0,0,0,0)",
	// 			},
	// 		},
	// 		{
	// 			id: "osm-mapnik",
	// 			type: "raster",
	// 			source: "osm-mapnik",
	// 		},
	// 	],
	// 	owner: "OpenStreetMap India",
	// },
};
