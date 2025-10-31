package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/blahaj-quest/ui"
	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohttp"
)

var (
	//go:embed public
	staticContent embed.FS
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	dataJSON, err := json.Marshal(data.Blahaj.Current)
	if err != nil {
		slog.Error("failed to get blahaj data json", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	emohttp.ServeOptimized(w, r, dataJSON, ".json", true)
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	html, err := ui.Render(r)
	if err != nil {
		slog.Error("failed to render: " + err.Error())
		http.Error(w, "failed to render", http.StatusInternalServerError)
		return
	}

	emohttp.ServeOptimized(w, r, []byte(html), ".html", true)
}

func main() {
	if config.IN_DEV {
		slog.Warn("in development mode")
		emohttp.DisableContentEncodingForHTML = true
		emohttp.PlausibleDisable = true
	}

	emohttp.PlausibleDomain = config.DOMAIN
	emohttp.PlausibleBaseURL = config.PLAUSIBLE_BASE_URL

	data.Init()

	err := goemo.InitSCSS(nil)
	if err != nil {
		slog.Error("failed to load scss transpiler: " + err.Error())
		os.Exit(1)
	}

	http.HandleFunc("GET /{$}", siteHandler)
	http.HandleFunc("GET /api/blahaj", apiHandler)
	http.HandleFunc("GET /notabot.gif", emohttp.HandleNotABotGif(
		func(r *http.Request) {
			emohttp.PlausibleEventFromNotABot(r)
		},
	))

	public, err := fs.Sub(staticContent, "public")
	if err != nil {
		slog.Error("failed to find public dir:" + err.Error())
		os.Exit(1)
	}
	http.HandleFunc("GET /{file...}", emohttp.FileServerOptimized(public))

	addr := fmt.Sprintf(":%s", config.PORT)

	slog.Info("starting http server at " + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("failed to start http server", "err", err)
	}
}
