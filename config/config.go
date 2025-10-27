package config

import (
	"os"
	"time"

	"github.com/makinori/blahaj-quest/util"
)

const (
	Title = "🔍 Blåhaj Quest"

	URL      = "https://blahaj.quest"
	ImageURL = URL + "/img/open-graph-image.jpg"

	Description = "Blåhaj loves you and needs you. Find them with this map so you can take good care of them ❤️"
	Keywords    = "blahaj, shark, ikea, stores, plush, plushie, stuffie, finder, quest, search"

	GitHubRepo = "makinori/blahaj-quest"
	GitHubURL  = "https://github.com/" + GitHubRepo

	CacheJSONPath   = "./cache.json"
	CacheExpireTime = time.Hour
)

func getEnv(name, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}

var (
	_, IN_DEV = os.LookupEnv("DEV")

	PORT = getEnv("PORT", "8080")

	COLOR         = "#3c8ea7"
	COLOR_DARKER  = util.MixHexColors(COLOR, "#000", 0.195)
	COLOR_LIGHTER = util.MixHexColors(COLOR, "#fff", 0.195)

	// MAP_STYLES = []struct {
	// 	Key  string `json:"key"`
	// 	Name string `json:"name"`
	// }{
	// 	{Key: "maptiler", Name: "MapTiler"},
	// 	{Key: "osm", Name: "OpenStreetMap"},
	// }

	// MAP_LAYERS = []struct {
	// 	Key  string `json:"key"`
	// 	Name string `json:"name"`
	// }{
	// 	{Key: "blahaj", Name: "Blåhaj"},
	// 	{Key: "heatmap", Name: "Heatmap"},
	// }
)
