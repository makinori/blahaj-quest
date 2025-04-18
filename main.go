package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/makinori/blahaj-quest/blahaj"
	"github.com/makinori/blahaj-quest/common"
	"github.com/makinori/blahaj-quest/ui"
	"github.com/makinori/blahaj-quest/ui/pages"
	"github.com/makinori/blahaj-quest/ui/render"

	"github.com/charmbracelet/log"
)

var (
	//go:embed public
	staticContent embed.FS

	_, isDev = os.LookupEnv("DEV")
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	dataJson, err := blahaj.GetBlahajDataJSON()

	if err != nil {
		log.Error("failed to get blahaj data json", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
}

type SiteData struct {
	URL         string
	Title       string
	Description string
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	data := blahaj.GetBlahajData()
	stars := common.GetGitHubStars()

	ctx := common.ChainContextValues(
		context.Background(),
		map[any]any{
			common.BlahajDataContextKey:  data,
			common.GitHubStarsContextKey: stars,
		},
	)

	html := render.Render(ctx, ui.Layout, pages.MainPage)

	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("GET /api/blahaj", apiHandler)
	http.HandleFunc("GET /{$}", siteHandler)

	if isDev {
		http.Handle("GET /", http.FileServerFS(os.DirFS("./public")))
	} else {
		public, err := fs.Sub(staticContent, "public")
		if err != nil {
			log.Fatal(err)
		}
		http.Handle("GET /", http.FileServerFS(public))
	}

	addr := fmt.Sprintf(":%d", 8080)

	log.Info("starting http server at " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("failed to start http server", "err", err)
	}
}
