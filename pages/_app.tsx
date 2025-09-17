import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import type { AppProps } from "next/app";
import Head from "next/head";
import Script from "next/script";
import { snPro } from "../fonts/fonts";
import { colorMix } from "../utils/utils";
import "./global.css";

// TODO: flashes on load, maybe ssr related

if (globalThis.localStorage != null) {
	globalThis.localStorage.setItem("chakra-ui-color-mode", "light");
}

const theme = extendTheme({
	initialColorMode: "light",
	useSystemColorMode: false,
	colors: {
		// material design pink
		// brand: {
		// 	50: "#fce4ec",
		// 	100: "#f8bbd0",
		// 	200: "#f48fb1",
		// 	300: "#f06292",
		// 	400: "#ec407a",
		// 	500: "#e91e63",
		// 	600: "#d81b60",
		// 	700: "#c2185b",
		// 	800: "#ad1457",
		// 	900: "#880e4f",
		// 	// a100: "#ff80ab",
		// 	// a200: "#ff4081",
		// 	// a400: "#f50057",
		// 	// a700: "#c51162",
		// },
		// tomorrow: "#1d1f21",
		// hexcorp: "#ff64ff",
		// hexcorpDark: "#231929",
		// justKindaDark: "#0f0f0f",
		// kofiBlue: {
		// 	50: "#13c3ff",
		// 	100: "#13c3ff",
		// 	200: "#13c3ff",
		// 	300: "#13c3ff",
		// 	400: "#13c3ff",
		// 	500: "#13c3ff",
		// 	600: colorMix("#13c3ff", "#111", 0.05),
		// 	700: colorMix("#13c3ff", "#111", 0.1),
		// 	800: colorMix("#13c3ff", "#111", 0.15),
		// 	900: colorMix("#13c3ff", "#111", 0.25),
		// },
	},
	components: {
		Heading: {
			baseStyle: {
				letterSpacing: "-0.02em",
				// fontWeight: "600",
			},
		},
		Link: {
			baseStyle: {
				color: "whiteAlpha.700",
				fontWeight: "700",
				_hover: {
					textDecoration: "none",
				},
			},
		},
	},
	styles: {
		global: {
			body: {
				bg: "justKindaDark",
				color: "white",
			},
		},
	},
	fonts: {
		heading: `${snPro.style.fontFamily}, sans-serif`,
		body: `${snPro.style.fontFamily}, sans-serif`,
	},
});

const title = "🔍 Blåhaj Quest";

const description =
	"Blåhaj loves you and needs you. Find them with this map so you can take good care of them ❤️";

const color = "#3c8ea7";

const url = "https://blahaj.quest/";
const imageUrl = "https://blahaj.quest/img/open-graph-image.jpg?2";

const keywords =
	"blahaj, shark, ikea, stores, trans, transgender, plush, plushie, stuffie, finder, quest, search";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<ChakraProvider theme={theme}>
			<Head>
				<meta charSet="utf-8" />
				<title>{title}</title>
				<meta
					name="viewport"
					content="width=device-width, initial-scale=0.7"
				/>
				<link
					rel="icon"
					type="image/png"
					sizes="32x32"
					href="favicon-32x32.png"
				/>
				<link
					rel="icon"
					type="image/png"
					sizes="16x16"
					href="favicon-16x16.png"
				/>
				<meta name="title" content={title} />
				<meta name="description" content={description} />
				<meta name="keywords" content={keywords} />
				<meta name="robots" content="index, follow" />
				<meta
					httpEquiv="Content-Type"
					content="text/html; charset=utf-8"
				/>

				<meta name="msapplication-TileColor" content={color} />
				{/* <meta
					name="msapplication-config"
					content="/assets/icons/browserconfig.xml"
				/> */}
				<meta name="theme-color" content={color} />
				<meta property="og:url" content={url} />
				<meta property="og:type" content="website" />
				<meta property="og:title" content={title} />
				<meta property="og:description" content={description} />
				<meta property="og:image" content={imageUrl} />
				<meta property="twitter:url" content={url} />
				<meta name="twitter:title" content={title} />
				<meta name="twitter:description" content={description} />
				<meta name="twitter:image" content={imageUrl} />
			</Head>
			<Script
				id="ithelpsme"
				defer
				data-domain="blahaj.quest"
				data-api="https://ithelpsme.hotmilk.space/api/event"
				dangerouslySetInnerHTML={{
					__html: `
						// prettier-ignore
						!function(){"use strict";var a=window.location,r=window.document,o=r.currentScript,l=o.getAttribute("data-api")||new URL(o.src).origin+"/api/event";function s(t,e){t&&console.warn("Ignoring Event: "+t),e&&e.callback&&e.callback()}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(a.hostname)||"file:"===a.protocol)return s("localhost",e);if(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)return s(null,e);try{if("true"===window.localStorage.plausible_ignore)return s("localStorage flag",e)}catch(t){}var n={},i=(n.n=t,n.u=a.href,n.d=o.getAttribute("data-domain"),n.r=r.referrer||null,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props),new XMLHttpRequest);i.open("POST",l,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&e&&e.callback&&e.callback()}}var e=window.plausible&&window.plausible.q||[];window.plausible=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function p(){n!==a.pathname&&(n=a.pathname,t("pageview"))}var c,w=window.history;w.pushState&&(c=w.pushState,w.pushState=function(){c.apply(this,arguments),p()},window.addEventListener("popstate",p)),"prerender"===r.visibilityState?r.addEventListener("visibilitychange",function(){n||"visible"!==r.visibilityState||p()}):p()}();
					`,
				}}
			></Script>
			<Component {...pageProps} />
		</ChakraProvider>
	);
}
