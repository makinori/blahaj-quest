package data

type blahajCountry struct {
	CountryCode      string
	LanguageCode     string
	ItemCode         string
	AdditionalStores []IkeaStore
}

func c(
	countryCode string, languageCode string, itemCode string,
	additionalStores ...IkeaStore,
) blahajCountry {
	return blahajCountry{
		CountryCode:      countryCode,
		LanguageCode:     languageCode,
		ItemCode:         itemCode,
		AdditionalStores: additionalStores,
	}
}

// find stores with
// https://www.ikea.com/us/en/meta-data/informera/stores-detailed.json

var BlahajDatabase = []blahajCountry{
	// --------
	// americas
	// --------

	// north america
	c("us", "en", "90373590"),
	c("ca", "en", "90373590"),
	c("mx", "es", "90373590"), // mexico
	c("cl", "es", "30373588"), // chile
	// dominica, regional site
	// pureto rico, regional site

	// ------
	// europe
	// ------

	// https://www.worldometers.info/geography/how-many-countries-in-europe/
	// c("ru", "ru", "30373588"), // russia
	c("de", "de", "30373588"), // germany
	c("gb", "en", "30373588"), // united kingdom
	c("fr", "fr", "30373588"), // france
	c("it", "it", "30373588"), // italy
	c("es", "es", "30373588"), // spain
	c("pl", "pl", "30373588"), // poland
	// should add a custom closed store but for now its closed
	// c("ua", "uk", "30373588"), // ukraine
	c("ro", "ro", "30373588"), // romania
	c("nl", "nl", "30373588"), // netherlands
	c("be", "nl", "30373588"), // belgium
	c("se", "sv", "30373588"), // sweden
	c("cz", "cs", "30373588"), // czech republic
	// greece has a different api
	c("pt", "pt", "30373588"), // portugal
	c("hu", "hu", "30373588"), // hungary
	// belarus has no ikea
	c("at", "de", "30373588"), // austria
	c("ch", "en", "30373588"), // switzerland
	c("rs", "sr", "30373588"), // serbia
	// bulgaria has a different api
	c("dk", "da", "30373588"), // denmark
	c("sk", "sk", "30373588"), // slovakia
	c("fi", "fi", "30373588"), // finland
	c("no", "no", "30373588"), // norway
	c("ie", "en", "30373588"), // ireland
	c("hr", "hr", "30373588"), // croatia
	// moldova has like fake ikea what
	// bosnia and herzegovina
	// albania
	// lithuania has a different api
	c("si", "sl", "30373588"), // slovenia
	// north macedonia has no ikea
	// latvia has a different api
	// c(
	// 	"ee",
	// 	"en",
	// 	"30373588",
	// 	IkeaStore{
	// 		ID:   "648",
	// 		Name: "IKEA Tallinn",
	// 		Lat:  "59.338481",
	// 		Lng:  "24.827531",
	// 	},
	// ), // estonia
	// luxembourg has no ikea yet, until 2025
	// montenegro has another fake ikea
	// malta
	// iceland has a different api
	// andorra has no ikea
	// liechtenstein has no ikea
	// monaco has no ikea
	// san marino has no ikea
	// holy see has no ikea

	// ----
	// asia
	// ----

	// afghanistan -
	// armenia -
	// azerbaijan -
	c("bh", "ar", "30373588"), // bahrain
	// bangladesh -
	// bhutan -
	// brunei -
	// burma/myanmar -
	// cambodia -
	c("cn", "zh", "10373589"), // china, redirected to local site, but same api
	// taiwan, regional site
	// hongkong and macau, regional site
	// cyprus, different api?
	// georgia -
	c("in", "en", "10373589"), // india
	// indonesia, regional? site
	// iran -
	// iraq -
	c("il", "he", "30373588"), // israel
	c("jp", "ja", "10373589"), // japan
	c("jo", "ar", "30373588"), // jordan
	// kazakhstan -
	c("kw", "ar", "30373588"), // kuwait
	// kyrgyzstan -
	// laos -
	// lebanon -
	c("my", "ms", "10373589"), // malaysia
	// maldives -
	// mongolia -
	// nepal -
	// north korea -
	// oman 2022
	// pakistan -
	// palestine -
	c("ph", "en", "10373589"), // philippines
	c("qa", "ar", "30373588"), // qatar
	c("sa", "ar", "30373588"), // saudi arabia
	c("sg", "en", "10373589"), // singapore
	c("kr", "ko", "10373589"), // south korea
	// sri lanka -
	// syria -
	// tajikstan -
	c("th", "th", "10373589"), // thailand
	// east timor -
	// turkey, regional? site
	// turkmenistan -
	c("ae", "ar", "30373588"), // united arab emirates
	// uzbekistan -
	// vietnam 2025
	// yemen -

	// ------
	// africa
	// ------

	c("eg", "ar", "30373588"), // egypt
	c("ma", "ar", "30373588"), // morocco
	// only 2 countries as of 2022-05-01

	// -------
	// oceania
	// -------

	c("au", "en", "10373589"), // australia
	c("nz", "en", "10373589"), // new zealand
	// only 2 countries as of 2025-01-03
}
