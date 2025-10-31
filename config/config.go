package config

import (
	"os"
	"time"

	"github.com/makinori/blahaj-quest/util"
)

const (
	Title = "🔍 Blåhaj Quest"

	URL       = "https://blahaj.quest"
	IMAGE_URL = URL + "/img/open-graph-image.jpg"

	DESCRIPTION = "Blåhaj loves you and needs you. Find them with this map so you can take good care of them ❤️"
	KEYWORDS    = "blahaj, shark, ikea, stores, plush, plushie, stuffie, finder, quest, search"

	GITHUB_REPO = "makinori/blahaj-quest"
	GITHUB_URL  = "https://github.com/" + GITHUB_REPO

	CACHE_JSON_PATH   = "./cache.json"
	CACHE_EXPIRE_TIME = time.Hour
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
	COLOR_DARKER  = util.MixHexColors(COLOR, "#000", 0.2)
	COLOR_LIGHTER = util.MixHexColors(COLOR, "#fff", 0.1)
)
