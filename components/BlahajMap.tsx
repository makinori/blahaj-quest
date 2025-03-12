import { Box, BoxProps } from "@chakra-ui/react";
import { Map, NavigationControl, Popup } from "maplibre-gl";
import "maplibre-gl/dist/maplibre-gl.css";
import { useCallback, useEffect, useRef } from "react";
import { BlahajData } from "../lib/get-blahaj";
import { MapStyleConfigs, useMapSettings } from "../settings/map-settings";
import styles from "./blahaj-map.module.css";

const mapLoadImage = (map: Map, url: string): Promise<HTMLImageElement> =>
	new Promise((resolve, reject) => {
		map.loadImage(url, function (error, image) {
			if (error) return reject(error);
			return resolve(image as HTMLImageElement);
		});
	});

const markerIconsSize = 2;
const markerIcons = {
	opaque: ["marker1", "marker2"],
	faded: ["marker1-faded", "marker2-faded"],
};

export default function BlahajMap(props: {
	boxProps?: BoxProps;
	blahajData: BlahajData;
}) {
	const mapRef = useRef<Map | null>(null);
	const mapContainer = useRef<HTMLDivElement>(null);

	const style = useMapSettings(s => s.style);
	const blahajLayer = useMapSettings(s => s.blahajLayer);
	const heatmapLayer = useMapSettings(s => s.heatmapLayer);

	const onStyleLoad = useCallback(
		async (blahajLayer: boolean, heatmapLayer: boolean) => {
			const map = mapRef.current;
			if (map == null) return;

			// trigger resize maybe helps?
			map.resize();

			// when change style, sources and layers get reset
			// transformStyle below handles this now.
			// just exit early if we already have "blahaj" source

			if (map.getSource("blahaj")) {
				return;
			}

			map.addSource("blahaj", {
				type: "geojson",
				data: {
					type: "FeatureCollection",
					features: props.blahajData.data.map(store => ({
						type: "Feature",
						properties: {
							icon: (store.quantity == 0
								? markerIcons.faded
								: markerIcons.opaque)[
								Math.floor(Math.random() * markerIconsSize)
							],
							weight: store.quantity / 32,
							description: `
${store.name}
<br />
<b>${store.quantity} blåhaj${store.quantity == 1 ? "" : "ar"}</b>
<br />
<a href="https://www.ikea.com/${store.countryCode}/${
								store.languageCode
							}/search/?q=blahaj">
See more →
</a>
							`.trim(),
						},
						geometry: {
							type: "Point",
							coordinates: [store.lng, store.lat],
						},
					})),
				},
			});

			map.addLayer({
				id: "heatmap",
				source: "blahaj",
				type: "heatmap",
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

			map.setLayoutProperty(
				"heatmap",
				"visibility",
				heatmapLayer ? "visible" : "none",
			);

			map.addLayer({
				id: "blahaj",
				source: "blahaj",
				type: "symbol",
				layout: {
					"icon-image": ["get", "icon"],
					"icon-size": 1,
					"icon-overlap": "always",
				},
			});

			map.setLayoutProperty(
				"blahaj",
				"visibility",
				blahajLayer ? "visible" : "none",
			);
		},
		[],
	);

	useEffect(() => {
		if (mapContainer.current == null) return;
		if (mapRef.current != null) return;

		// const protocol = new pmtiles.Protocol();
		// MapLibreGL.addProtocol("pmtiles", protocol.tile);

		const map = new Map({
			container: mapContainer.current,
			style: MapStyleConfigs[style],
			center: [15, 20],
			zoom: 1.8,
		});

		mapRef.current = map;

		map.addControl(new NavigationControl());

		map.dragRotate.disable();
		map.touchZoomRotate.disableRotation();

		for (const name of [...markerIcons.opaque, ...markerIcons.faded]) {
			mapLoadImage(map, "img/marker-icons/48/" + name + ".png")
				.then(image => {
					map.addImage(name, image);
				})
				.catch(error => {
					console.error(error);
				});
		}

		map.on("mouseenter", "markers", () => {
			map.getCanvas().style.cursor = "pointer";
		});

		map.on("mouseleave", "markers", () => {
			map.getCanvas().style.cursor = "";
		});

		map.on("click", "blahaj", e => {
			const feature = e.features[0] as any;
			const coordinates = feature.geometry.coordinates.slice();
			const description = feature.properties.description;
			const popup = new Popup()
				.setLngLat(coordinates)
				.addTo(map)
				.setHTML(description);

			try {
				const container = popup._container;

				const content = container.querySelector(
					".maplibregl-popup-content",
				);

				content.className += " " + styles["blahaj-popup-content"];

				const closeButton = content.querySelector(
					".maplibregl-popup-close-button",
				);

				closeButton.className +=
					" " + styles["blahaj-popup-close-button"];
			} catch (error) {
				console.error(error);
			}
		});

		map.once("style.load", () => {
			onStyleLoad(blahajLayer, heatmapLayer);
		});
	}, []);

	useEffect(() => {
		try {
			mapRef.current.setLayoutProperty(
				"blahaj",
				"visibility",
				blahajLayer ? "visible" : "none",
			);
		} catch (error) {}
	}, [blahajLayer]);

	useEffect(() => {
		try {
			mapRef.current.setLayoutProperty(
				"heatmap",
				"visibility",
				heatmapLayer ? "visible" : "none",
			);
		} catch (error) {}
	}, [heatmapLayer]);

	useEffect(() => {
		try {
			if (!mapRef.current.isStyleLoaded()) return;
			mapRef.current.setStyle(MapStyleConfigs[style], {
				transformStyle: (prev, next) => {
					if (next.sources["blahaj"] == null) {
						next.sources["blahaj"] = prev.sources["blahaj"];
					}

					if (!next.layers.find(l => l.id == "heatmap")) {
						next.layers.push(
							prev.layers.find(l => l.id == "heatmap"),
						);
					}

					if (!next.layers.find(l => l.id == "blahaj")) {
						next.layers.push(
							prev.layers.find(l => l.id == "blahaj"),
						);
					}

					return next;
				},
			});
			onStyleLoad(blahajLayer, heatmapLayer);
		} catch (error) {}
	}, [style]);

	return <Box ref={mapContainer} {...props.boxProps} />;
}
