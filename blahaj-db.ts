// country code, language code, product id, additional stores

// find stores with
// https://www.ikea.com/us/en/meta-data/informera/stores-detailed.json

export interface BlahajDbStore {
	id: string;
	name: string;
	lat: string;
	lng: string;
}

export type BlahajDbCountry = [string, string, string, BlahajDbStore[]?];

export const BLAHAJ_DB: { [continent: string]: BlahajDbCountry[] } = {
	americas: [
		// north america
		["us", "en", "90373590"],
		["ca", "en", "90373590"],
		["mx", "es", "90373590"], // mexico
		["cl", "es", "30373588"], // chile
		// dominica, regional site
		// pureto rico, regional site
	],
	europe: [
		// https://www.worldometers.info/geography/how-many-countries-in-europe/
		["ru", "ru", "30373588"], // russia
		["de", "de", "30373588"], // germany
		["gb", "en", "30373588"], // united kingdom
		["fr", "fr", "30373588"], // france
		["it", "it", "30373588"], // italy
		["es", "es", "30373588"], // spain
		["pl", "pl", "30373588"], // poland
		// should add a custom closed store but for now its closed
		// ["ua", "uk", "30373588"], // ukraine
		["ro", "ro", "30373588"], // romania
		["nl", "nl", "30373588"], // netherlands
		["be", "nl", "30373588"], // belgium
		["se", "sv", "30373588"], // sweden
		["cz", "cs", "30373588"], // czech republic
		// greece has a different api
		["pt", "pt", "30373588"], // portugal
		["hu", "hu", "30373588"], // hungary
		// belarus has no ikea
		["at", "de", "30373588"], // austria
		["ch", "en", "30373588"], // switzerland
		["rs", "sr", "30373588"], // serbia
		// bulgaria has a different api
		["dk", "da", "30373588"], // denmark
		["sk", "sk", "30373588"], // slovakia
		["fi", "fi", "30373588"], // finland
		["no", "no", "30373588"], // norway
		["ie", "en", "30373588"], // ireland
		["hr", "hr", "30373588"], // croatia
		// moldova has like fake ikea what
		// bosnia and herzegovina
		// albania
		// lithuania has a different api
		["si", "sl", "30373588"], // slovenia
		// north macedonia has no ikea
		// latvia has a different api
		[
			"ee",
			"en",
			"30373588",
			[
				{
					id: "648",
					name: "IKEA Tallinn",
					lat: "59.338481",
					lng: "24.827531",
				},
			],
		], // estonia
		// luxembourg has no ikea yet, until 2025
		// montenegro has another fake ikea
		// malta
		// iceland has a different api
		// andorra has no ikea
		// liechtenstein has no ikea
		// monaco has no ikea
		// san marino has no ikea
		// holy see has no ikea
	],
	asia: [
		// afghanistan -
		// armenia -
		// azerbaijan -
		["bh", "ar", "30373588"], // bahrain
		// bangladesh -
		// bhutan -
		// brunei -
		// burma/myanmar -
		// cambodia -
		["cn", "zh", "10373589"], // china, redirected to local site, but same api
		// taiwan, regional site
		// hongkong and macau, regional site
		// cyprus, different api?
		// georgia -
		["in", "en", "10373589"], // india
		// indonesia, regional? site
		// iran -
		// iraq -
		["il", "he", "30373588"], // israel
		["jp", "ja", "10373589"], // japan
		["jo", "ar", "30373588"], // jordan
		// kazakhstan -
		["kw", "ar", "30373588"], // kuwait
		// kyrgyzstan -
		// laos -
		// lebanon -
		["my", "ms", "10373589"], // malaysia
		// maldives -
		// mongolia -
		// nepal -
		// north korea -
		// oman 2022
		// pakistan -
		// palestine -
		["ph", "en", "10373589"], // philippines
		["qa", "ar", "30373588"], // qatar
		["sa", "ar", "30373588"], // saudi arabia
		["sg", "en", "10373589"], // singapore
		["kr", "ko", "10373589"], // south korea
		// sri lanka -
		// syria -
		// tajikstan -
		["th", "th", "10373589"], // thailand
		// east timor -
		// turkey, regional? site
		// turkmenistan -
		["ae", "ar", "30373588"], // united arab emirates
		// uzbekistan -
		// vietnam 2025
		// yemen -
	],
	africa: [
		["eg", "ar", "30373588"], // egypt
		["ma", "ar", "30373588"], // morocco
		// only 2 country as of 2022-05-01
	],
	oceania: [
		["au", "en", "10373589"], // australia
		// new zealand, 2022
		// only 1 country and 1 planned as of 2022-05-01
	],
};
