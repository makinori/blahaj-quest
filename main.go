package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/blahaj-quest/ui"
	"github.com/makinori/goemo"
)

var (
	//go:embed public
	staticContent embed.FS
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	dataJson, err := data.GetBlahajDataJSON()

	if err != nil {
		slog.Error("failed to get blahaj data json", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	html, err := ui.Render()

	if err != nil {
		slog.Error("failed to render: " + err.Error())
		http.Error(w, "failed to render", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

func main() {
	if config.InDev {
		slog.Warn("in development mode")
	}

	err := goemo.InitSCSS()
	if err != nil {
		slog.Error("failed to load scss transpiler: " + err.Error())
		os.Exit(1)
	}

	http.HandleFunc("GET /api/blahaj", apiHandler)
	http.HandleFunc("GET /{$}", siteHandler)

	if config.InDev {
		http.Handle("GET /", http.FileServerFS(os.DirFS("./public")))
	} else {
		public, err := fs.Sub(staticContent, "public")
		if err != nil {
			slog.Error("failed to find public dir:" + err.Error())
			os.Exit(1)
		}
		http.Handle("GET /", http.FileServerFS(public))
	}

	addr := fmt.Sprintf(":%s", config.Port)

	slog.Info("starting http server at " + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("failed to start http server", "err", err)
	}
}
