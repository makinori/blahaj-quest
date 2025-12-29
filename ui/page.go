package ui

import (
	"context"
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"github.com/makinori/foxlib/foxcss"
	"github.com/makinori/foxlib/foxhtml"
	"github.com/makinori/foxlib/foxjs"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed font.css
	fontCSS string
	//go:embed page.css
	pageCSS string

	//go:embed page.js
	pageJS string
)

func init() {
	fontCSS = foxcss.MustMinify(fontCSS)
	pageCSS = foxcss.MustMinify(pageCSS)
	pageJS = foxjs.MustMinify(pageJS)
}

func Render(r *http.Request) (string, error) {
	ctx := foxcss.InitContext(context.Background(), "")

	err := foxcss.UseWords(ctx, hajWords, time.Now().Format(time.DateOnly))
	if err != nil {
		return "", err
	}

	// TODO: compress this better
	blahajData, err := json.Marshal(data.Blahaj.Current)
	if err != nil {
		slog.Error("failed to marshal blahaj data", "err", err)
	}

	body := Body(
		foxhtml.VStack(ctx,
			foxhtml.StackCSS(`
				gap: 0;
				width: 100vw;
				height: 100vh;
			`),
			BlahajHeader(ctx, r),
			BlahajMap(ctx),
		),
		Script(Raw(
			"const blahajData = "+string(blahajData)+";"+pageJS,
		)),
	)

	css := fontCSS + pageCSS + foxcss.GetPageCSS(ctx)
	// os.WriteFile("test.css", []byte(css), 0644)

	head := Head(
		Meta(Charset("utf-8")),
		TitleEl(Text(config.TITLE)),
		Meta(
			Name("viewport"),
			Content("width=device-width, initial-scale=0.8"),
		),
		Link(
			Rel("icon"), Type("image/png"), Attr("sizes", "16x16"),
			Href("/favicon-16x16.png"),
		),
		Link(
			Rel("icon"), Type("image/png"), Attr("sizes", "32x32"),
			Href("/favicon-32x32.png"),
		),
		// TODO: check if meta tags are correct
		Meta(Name("title"), Content(config.TITLE)),
		Meta(Name("description"), Content(config.DESCRIPTION)),
		Meta(Name("keywords"), Content(config.KEYWORDS)),
		Meta(Name("robots"), Content("index, follow")),
		Meta(Attr("http-equiv", "text/html; charset=utf-8")),
		Meta(Name("msapplication-TileColor"), Content(config.COLOR)),
		Meta(Name("theme-color"), Content(config.COLOR)),
		Meta(Attr("property", "og:url"), Content(config.URL)),
		Meta(Attr("property", "og:type"), Content("website")),
		Meta(Attr("property", "og:title"), Content(config.TITLE)),
		Meta(Attr("property", "og:description"), Content(config.DESCRIPTION)),
		Meta(Attr("property", "og:image"), Content(config.IMAGE_URL)),
		Meta(Name("twitter:title"), Content(config.TITLE)),
		Meta(Name("twitter:description"), Content(config.DESCRIPTION)),
		Meta(Name("twitter:image"), Content(config.IMAGE_URL)),
		Script(Src("/js/maplibre-gl.js")),
		Link(Href("/css/maplibre-gl.css"), Rel("stylesheet")),
		StyleEl(Raw(css)),
	)

	page := Doctype(
		HTML(
			head,
			body,
		),
	)

	return Group{page}.String(), nil
}
