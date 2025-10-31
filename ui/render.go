package ui

import (
	"context"
	_ "embed"
	"time"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/goemo"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed global.scss
var globalSCSS string

func Render() (string, error) {
	ctx := goemo.InitContext(context.Background())

	ctx = goemo.UseWords(
		ctx, goemo.AnimalWords, time.Now().Format(time.DateOnly),
	)

	body := Body(
		BlahajHeader(ctx),
		BlahajMap(),
	)

	pageSCSS := goemo.GetPageSCSS(ctx)

	pageCSS, err := goemo.RenderSCSS(globalSCSS + "\n" + pageSCSS)
	if err != nil {
		return "", err
	}

	head := Head(
		Meta(Charset("utf-8")),
		TitleEl(Text(config.Title)),
		Meta(
			Name("viewport"),
			Content("width=device-width, initial-scale=0.7"),
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
		Meta(Name("title"), Content(config.Title)),
		Meta(Name("description"), Content(config.DESCRIPTION)),
		Meta(Name("keywords"), Content(config.KEYWORDS)),
		Meta(Name("robots"), Content("index, follow")),
		Meta(Attr("http-equiv", "text/html; charset=utf-8")),
		Meta(Name("msapplication-TileColor"), Content(config.COLOR)),
		Meta(Name("theme-color"), Content(config.COLOR)),
		Meta(Attr("property", "og:url"), Content(config.URL)),
		Meta(Attr("property", "og:type"), Content("website")),
		Meta(Attr("property", "og:title"), Content(config.Title)),
		Meta(Attr("property", "og:description"), Content(config.DESCRIPTION)),
		Meta(Attr("property", "og:image"), Content(config.IMAGE_URL)),
		Meta(Name("twitter:title"), Content(config.Title)),
		Meta(Name("twitter:description"), Content(config.DESCRIPTION)),
		Meta(Name("twitter:image"), Content(config.IMAGE_URL)),
		StyleEl(Raw(pageCSS)),
	)

	page := Doctype(
		HTML(
			head,
			body,
		),
	)

	return Group{page}.String(), nil
}
