import {
	Box,
	Checkbox,
	HStack,
	Radio,
	RadioGroup,
	Text,
	VStack,
} from "@chakra-ui/react";
import {
	MapStyle,
	MapStyleConfigs,
	useMapSettings,
} from "../settings/map-settings";

export default function MapSettings() {
	const settings = useMapSettings();

	const mapStylesKeys = Object.keys(MapStyleConfigs);
	let mapStyleButtons: JSX.Element[] = [];

	for (const key of mapStylesKeys) {
		mapStyleButtons.push(
			<Radio value={key}>
				<Text fontSize={14}>{key}</Text>
			</Radio>,
		);
		if (key != mapStylesKeys[mapStylesKeys.length - 1]) {
			mapStyleButtons.push(
				<Box borderX={"solid 1px rgba(0,0,0,0.1)"} h="32px" />,
			);
		}
	}

	return (
		<>
			<Box
				bg={"white"}
				p={"0px 12px"}
				boxShadow={"lg"}
				borderRadius={12}
				color={"black"}
				alignItems={"start"}
				mb={2}
			>
				<RadioGroup
					value={settings.style}
					onChange={e => settings.setStyle(e as MapStyle)}
				>
					<HStack spacing={3}>{mapStyleButtons}</HStack>
				</RadioGroup>
			</Box>
			<VStack
				display={"inline-flex"}
				bg={"white"}
				p={"6px 12px"}
				boxShadow={"lg"}
				borderRadius={12}
				color={"black"}
				alignItems={"start"}
				spacing={0}
				fontWeight={500}
			>
				<Checkbox
					defaultChecked={settings.blahajLayer}
					isChecked={settings.blahajLayer}
					onChange={e => settings.setBlahajLayer(e.target.checked)}
				>
					<Text fontSize={14}>blåhaj</Text>
				</Checkbox>
				<Checkbox
					defaultChecked={settings.heatmapLayer}
					isChecked={settings.heatmapLayer}
					onChange={e => settings.setHeatmapLayer(e.target.checked)}
				>
					<Text fontSize={14}>heatmap</Text>
				</Checkbox>
			</VStack>
		</>
	);
}
