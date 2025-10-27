package data

import "github.com/makinori/goemo/emocache"

func Init() {
	emocache.Init("cache", []emocache.DataInterface{
		&Blahaj, &GitHubStars,
	})
}
