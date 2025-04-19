package common

import "time"

const (
	ConfigTitle = "🔍 Blåhaj Quest"

	ConfigURL      = "https://blahaj.quest"
	ConfigImageURL = ConfigURL + "/img/open-graph-image.jpg"

	ConfigDescription = "Blåhaj loves you and needs you. Find them with this map so you can take good care of them ❤️"
	ConfigKeywords    = "blahaj, shark, ikea, stores, plush, plushie, stuffie, finder, quest, search"

	ConfigGitHubRepo = "makinori/blahaj-quest"
	ConfigGitHubURL  = "https://github.com/" + ConfigGitHubRepo

	ConfigCacheJSONPath   = "./cache.json"
	ConfigCacheExpireTime = time.Hour
)

var (
	ConfigColor       = "#3c8ea7"
	ConfigColorDarker = MixHexColors(ConfigColor, "#000", 0.195)
	ConfigColorLigher = MixHexColors(ConfigColor, "#fff", 0.195)
)

var ConfigMapStyles = []struct {
	Key   string
	Title string
}{
	{Key: "maptiler", Title: "MapTiler"},
	{Key: "osm", Title: "OpenStreetMap"},
}
