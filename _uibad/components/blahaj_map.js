const map = document.getElementById("map");

console.log(map);

function onMapStyle(mapStyle) {
	console.log(mapStyle);
}

function onMapLayer(mapLayer, checked) {
	console.log(mapLayer, checked);
}

for (let i in MapStyles) {
	const { Key } = MapStyles[i];
	const el = document.getElementById("map-style-" + Key);
	el.checked = i == 0;
	el.addEventListener("change", e => {
		onMapStyle(Key);
	});
}

for (const { Key } of MapLayers) {
	const el = document.getElementById("map-layer-" + Key);
	el.checked = true;
	el.addEventListener("change", e => {
		onMapLayer(Key, e.target.checked);
	});
}
