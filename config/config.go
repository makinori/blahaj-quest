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
	_, InDev = os.LookupEnv("DEV")

	Port = getEnv("PORT", "8080")

	Color       = "#3c8ea7"
	ColorDarker = util.MixHexColors(Color, "#000", 0.195)
	ColorLigher = util.MixHexColors(Color, "#fff", 0.195)
)

var MapStyles = []struct {
	Key   string
	Title string
}{
	{Key: "maptiler", Title: "MapTiler"},
	{Key: "osm", Title: "OpenStreetMap"},
}

var MapLayers = []struct {
	Key   string
	Title string
}{
	{Key: "blahaj", Title: "Blåhaj"},
	{Key: "heatmap", Title: "Heatmap"},
}
