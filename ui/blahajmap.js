const el = document.currentScript.parentElement;
const mapEl = el.querySelector("#map");

const mapStyleEls = document.querySelectorAll("input[name=map-style]");
const mapLayerEls = document.querySelectorAll("input[name=map-layer]");

for (const i in mapStyleEls) {
	if (i == 0) {
		mapStyleEls[i].checked = true;
	} else {
		mapStyleEls[i].checked = false;
	}
}

for (const i in mapLayerEls) {
	mapLayerEls[i].checked = true;
}

const getInputEl = (els, name) => {
	for (const el of els) {
		if (el.value == name) {
			return el;
		}
	}
	return null;
};

const map = new maplibregl.Map({
	container: mapEl,
	center: [15, 20],
	zoom: 1.8,
});

map.addControl(
	new maplibregl.NavigationControl({
		showCompass: false,
	}),
);

map.dragRotate.disable();
map.touchZoomRotate.disableRotation();

const markerIconsSize = 2;
const markerIcons = {
	opaque: ["marker1", "marker2"],
	faded: ["marker1-faded", "marker2-faded"],
};

for (const name of [...markerIcons.opaque, ...markerIcons.faded]) {
	// TODO: promise.all
	try {
		const image = await map.loadImage(
			"/img/marker-icons/48/" + name + ".png",
		);
		map.addImage(name, image.data);
	} catch (error) {
		console.error(error);
	}
}

map.on("mouseenter", "markers", () => {
	map.getCanvas().style.cursor = "pointer";
});

map.on("mouseleave", "markers", () => {
	map.getCanvas().style.cursor = "";
});

map.on("click", "blahaj", e => {
	const feature = e.features[0];
	const coordinates = feature.geometry.coordinates.slice();
	const description = feature.properties.description;
	new maplibregl.Popup()
		.setLngLat(coordinates)
		.addTo(map)
		.setHTML(description);
});

// ${store.quantity == 1 ? "" : "ar"}
const makeDescription = store => `${store.name}
<br />
<b>${store.quantity} blåhaj</b>
<br />
<a href="https://www.ikea.com/${store.countryCode}/${store.languageCode}/search/?q=blahaj">
See more →
</a>`;

const blahajSource = {
	type: "geojson",
	data: {
		type: "FeatureCollection",
		features: blahajData.map(store => ({
			type: "Feature",
			properties: {
				icon: (store.quantity == 0
					? markerIcons.faded
					: markerIcons.opaque)[
					Math.floor(Math.random() * markerIconsSize)
				],
				weight: store.quantity / 32,
				description: makeDescription(store),
			},
			geometry: {
				type: "Point",
				coordinates: [store.lng, store.lat],
			},
		})),
	},
};

const ensureLayer = (style, layer) => {
	const layerIndex = style.layers.findIndex(l => l.id == layer.id);
	if (layerIndex == -1) {
		style.layers.push(layer);
		return;
	}

	style.layers.splice(layerIndex, 1);
	style.layers.push(layer);
};

const getLayerVisibility = name =>
	getInputEl(mapLayerEls, name).checked ? "visible" : "none";

const transformStyle = (prev, style) => {
	style.sources.blahaj = blahajSource;

	ensureLayer(style, {
		id: "blahaj-heatmap",
		source: "blahaj",
		type: "heatmap",
		layout: {
			visibility: getLayerVisibility("blahaj-heatmap"),
		},
		paint: {
			"heatmap-weight": {
				property: "weight",
				type: "identity",
			},
			// "heatmap-radius": 60,
			"heatmap-radius": [
				"interpolate",
				["linear"],
				["zoom"],
				0,
				20,
				100,
				1000,
			],
			"heatmap-opacity": 0.5,
		},
	});

	ensureLayer(style, {
		id: "blahaj",
		source: "blahaj",
		type: "symbol",
		layout: {
			visibility: getLayerVisibility("blahaj"),
			"icon-image": ["get", "icon"],
			"icon-size": 1,
			"icon-overlap": "always",
		},
	});

	return style;
};

const mapStyles = {
	openfreemap: "https://tiles.openfreemap.org/styles/liberty",
	openstreetmap: {
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
	},
};

map.setStyle(mapStyles[mapStyleEls[0].value], {
	transformStyle,
});

for (const mapStyleEl of mapStyleEls) {
	mapStyleEl.addEventListener("change", () => {
		map.setStyle(mapStyles[mapStyleEl.value], {
			transformStyle,
		});
	});
}

for (const mapLayerEl of mapLayerEls) {
	mapLayerEl.addEventListener("change", () => {
		map.setLayoutProperty(
			mapLayerEl.value,
			"visibility",
			mapLayerEl.checked ? "visible" : "none",
		);
	});
}
