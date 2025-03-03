import axios from "axios";
import * as fs from "fs/promises";
import pLimit from "p-limit";
import * as path from "path";
import { BLAHAJ_DB, BlahajDbCountry, BlahajDbStore } from "../blahaj-db";

const blahajDataCachePath = path.resolve(__dirname, "blahaj.json");
const blahajDataExpireTime = 1000 * 60 * 60; // 1 hour

export interface BlahajStore {
	quantity: number;
	name: string;
	lat: string;
	lng: string;
	countryCode: string;
	languageCode: string;
	itemCode: string;
}

export interface BlahajData {
	updated: string;
	data: BlahajStore[];
}

async function getStock(country: BlahajDbCountry): Promise<BlahajStore[]> {
	const [countryCode, languageCode, itemCode, additionalStores] = country;

	try {
		let stores: BlahajDbStore[] = [];

		// add from additional
		if (additionalStores != null) {
			stores = additionalStores;
		}

		// try fetching stores from url
		try {
			const foundStores = await axios<BlahajDbStore[]>({
				url: `https://www.ikea.com/${countryCode}/${languageCode}/meta-data/informera/stores-detailed.json`,
			});
			stores = foundStores.data;
		} catch (error) {}

		const stock = await axios<{
			availabilities: {
				buyingOption: {
					cashCarry: { availability: { quantity: number } };
				};
				classUnitKey: {
					classUnitCode: string;
				};
			}[];
		}>({
			url: `https://api.ingka.ikea.com/cia/availabilities/ru/${countryCode}`,
			headers: {
				Accept: "application/json;version=2",
				Referer: "https://www.ikea.com/",
				// "X-Client-Id": "b6c117e5-ae61-4ef5-b4cc-e0b1e37f0631",
				"X-Client-Id": "ef382663-a2a5-40d4-8afe-f0634821c0ed",
			},
			params: {
				itemNos: itemCode,
				// expand: "StoresList,Restocks,SalesLocations",
				expand: "StoresList",
			},
		});

		return stock.data.availabilities
			.map(storeAvail => {
				const quantity =
					storeAvail?.buyingOption?.cashCarry?.availability?.quantity;
				if (quantity == null) return null;

				const storeId = storeAvail?.classUnitKey?.classUnitCode;
				const store = stores.find(store => store.id == storeId);
				if (store == null) return null;

				return {
					quantity,
					name: store.name,
					lat: store.lat,
					lng: store.lng,
					countryCode,
					languageCode,
					itemCode,
				} as BlahajStore;
			})
			.filter(store => store != null) as BlahajStore[];
	} catch (error) {
		console.error(countryCode + "-" + languageCode + " failed");
		console.error(error);
		return [];
	}
}

async function setCache(data: BlahajData) {
	await fs.writeFile(blahajDataCachePath, JSON.stringify(data));
}

async function getCache(): Promise<BlahajData | null> {
	try {
		const json = await fs.readFile(blahajDataCachePath, "utf8");
		return JSON.parse(json);
	} catch (error) {
		return null;
	}
}

export async function getBlahajData(): Promise<BlahajData> {
	// get cache if available and not expired

	const blahajDataCache = await getCache();

	if (
		blahajDataCache != null &&
		Date.now() - new Date(blahajDataCache?.updated).getTime() <
			blahajDataExpireTime
	) {
		return blahajDataCache;
	}

	// get fresh data

	const countries = Object.values(BLAHAJ_DB).flat();
	// const countries = [
	// 	...BLAHAJ_DB.americas,
	// 	...BLAHAJ_DB.europe,
	// 	...BLAHAJ_DB.asia,
	// 	...BLAHAJ_DB.africa,
	// 	...BLAHAJ_DB.oceania,
	// ];

	const limit = pLimit(10); // requests at a time
	const stockResponse = await Promise.all(
		countries.map(info => limit(() => getStock(info))),
	);

	const stores = stockResponse.flat();
	console.log(`Fetched data from ${stores.length} stores`);

	const blahajData = {
		updated: new Date().toISOString(),
		data: stores,
	};

	await setCache(blahajData);

	return blahajData;
}
