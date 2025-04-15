import {
	Box,
	Button,
	HStack,
	Heading,
	Icon,
	Image,
	Link,
	Text,
	VStack,
	useBreakpointValue,
} from "@chakra-ui/react";
import axios from "axios";
import * as dateFns from "date-fns";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { FaCode } from "react-icons/fa6";
import BlahajMap from "../components/BlahajMap";
import { GitHubIcon } from "../components/GitHubIcon";
import MapSettings from "../components/MapSettings";
import blahajImage from "../images/full-flipped.png";
import { BlahajData, getBlahajData } from "../lib/get-blahaj";
import { apiCache } from "../utils/api-cache";

async function getGitHubStars() {
	return apiCache<string>(
		"github-stars",
		async () => {
			try {
				return String(
					(
						await axios(
							"https://api.github.com/repos/makinori/blahaj-quest",
						)
					).data.stargazers_count,
				);
			} catch (error) {
				return "?";
			}
		},
		1000 * 60 * 60, // 1 hour
	);
}

export const getServerSideProps = (async context => {
	return {
		props: {
			blahajData: await getBlahajData(),
			gitHubStars: await getGitHubStars(),
		},
	};
}) satisfies GetServerSideProps<{
	blahajData: BlahajData;
	gitHubStars: string;
}>;

export default function Home(
	props: InferGetServerSidePropsType<typeof getServerSideProps>,
) {
	const breakpoint = useBreakpointValue([0, 1, 2, 3]);

	const Logo = (
		<>
			<Image
				src={blahajImage.src}
				w={"195px"}
				mt={"-32px"}
				ml={"-96px"}
			></Image>
			<VStack spacing={0.5}>
				<Heading
					fontSize={24}
					lineHeight={0.8}
					fontWeight={800}
					mt={1.5}
				>
					blåhaj quest
				</Heading>
				<Link href="https://blahaj.quest">https://blahaj.quest</Link>
			</VStack>
		</>
	);

	const Stats = (
		<VStack spacing={0} ml={4} alignItems={"start"} mt={0}>
			<HStack spacing={4}>
				<Text fontWeight={500}>
					<b>
						{props.blahajData.data.length.toLocaleString("en-US")}{" "}
						stores
					</b>
					{" with "}
					<b>
						{props.blahajData.data
							.reduce((a, b) => a + b.quantity, 0)
							.toLocaleString("en-US")}{" "}
						blåhaj
					</b>
				</Text>
			</HStack>
			<Text mt={-1} fontWeight={500}>
				last updated:{" "}
				<b>
					{/* {new Date(props.blahajData.updated).toLocaleString([], {
						hour: "2-digit",
						minute: "2-digit",
					})} */}
					{dateFns.formatDistance(
						new Date(props.blahajData.updated),
						new Date(),
						{ addSuffix: true },
					)}
				</b>
			</Text>
		</VStack>
	);

	const Credits = (
		<HStack mt={-1.5} mr={4} spacing={4}>
			<VStack spacing={1.5}>
				{/* <HStack spacing={1}>
					<Text fontSize={"sm"}>
						Made by{" "}
						<Link href="https://maki.cafe" color={"white"}>
							Maki
						</Link>
					</Text>
					<Image src={transHeart.src} h={5} />
				</HStack>
				<Button
					as={"a"}
					size={"xs"}
					leftIcon={<KofiIcon />}
					colorScheme={"kofiBlue"}
					color={"white"}
					outline={"solid 2px rgba(255,255,255,0.5)"}
					href="https://ko-fi.com/maki_nori"
				>
					Support me
				</Button> */}
			</VStack>
			<VStack spacing={1}>
				<Link
					href="https://github.com/makinori/blahaj-quest"
					// mt={-1.5}
					display={"flex"}
					flexDir={"row"}
					alignItems={"center"}
					justifyContent={"center"}
					fontWeight={600}
					fontSize={"sm"}
					color={"white"}
				>
					<Icon as={FaCode} mr={1.5} />
					see code
				</Link>
				<Button
					as={"a"}
					size={"xs"}
					leftIcon={<GitHubIcon />}
					colorScheme={"gray"}
					color={"black"}
					outline={"solid 2px rgba(255,255,255,0.5)"}
					href="https://github.com/makinori/blahaj-quest"
					overflow={"hidden"}
				>
					Star
					<Box
						background={"white"}
						h={"100%"}
						mr={"-8px"}
						ml={"8px"}
						px={"8px"}
						display={"flex"}
						alignItems={"center"}
						justifyContent={"center"}
						borderLeft={"solid 1px var(--chakra-colors-gray-300)"}
					>
						{props.gitHubStars}
					</Box>
				</Button>
				{/* <Box mt={1}>
					<GitHubButton
						href="https://github.com/makinori/blahaj-quest"
						data-size="large"
						data-show-count="true"
						aria-label="Star makinori/blahaj-quest on GitHub"
					>
						Star
					</GitHubButton>
				</Box> */}
			</VStack>
		</HStack>
	);

	let InnerHeader = <></>;

	if (breakpoint > 2) {
		// wide
		InnerHeader = (
			<HStack h={"64px"} w={"100vw"} spacing={4} zIndex={999}>
				{Logo}
				{Stats}
				<Box flexGrow={1}></Box>
				{Credits}
			</HStack>
		);
	} else {
		// compact
		InnerHeader = (
			<>
				<HStack h={"64px"} w={"100vw"} spacing={4} zIndex={999}>
					{Logo}
					<Box flexGrow={1}></Box>
					{Credits}
				</HStack>
				<HStack
					h={"64px"}
					w={"100vw"}
					spacing={4}
					zIndex={999}
					justifyContent={"center"}
					background={"rgba(255,255,255,0.1)"}
				>
					{Stats}
				</HStack>
			</>
		);
	}

	return (
		<VStack flexDir={"column"} w={"100vw"} h={"100vh"} spacing={0}>
			<VStack spacing={0} bg={"#3C8EA7"} boxShadow={"lg"} zIndex={100}>
				{InnerHeader}
				<Box
					bg={"rgba(0,0,0,0.2)"}
					w={"100%"}
					h={"8px"}
					pos={"relative"}
				>
					<Box
						pos={"absolute"}
						top={"16px"}
						left={"8px"}
						zIndex={2000}
					>
						<MapSettings />
					</Box>
				</Box>
			</VStack>
			<BlahajMap
				blahajData={props.blahajData}
				boxProps={{ w: "100vw", flexGrow: 1 }}
			/>
		</VStack>
	);
}
