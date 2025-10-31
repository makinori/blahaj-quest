package data

import (
	"io"
	"net/http"

	"github.com/makinori/goemo/emocache"
)

func getItHelpsMe() (string, error) {
	res, err := http.Get("https://ithelpsme.hotmilk.space/js/script.js")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

var ItHelpsMe = emocache.Data[string]{
	Key:      "ithelpsme",
	CronSpec: "0 0 * * *", // start of every day
	Retrieve: getItHelpsMe,
}
