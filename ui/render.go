package ui

import (
	"context"
	_ "embed"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/goemo"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed global.scss
var globalSCSS string

func Render() (string, error) {
	ctx := goemo.InitContext(context.Background())

	pageSCSS := goemo.GetPageSCSS(ctx)

	pageCSS, err := goemo.RenderSCSS(globalSCSS + "\n" + pageSCSS)
	if err != nil {
		return "", err
	}

	head := Head(
		Meta(Charset("utf-8")),
		Title(config.Title),
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
		Meta(Name("description"), Content(config.Description)),
		Meta(Name("keywords"), Content(config.Keywords)),
		Meta(Name("robots"), Content("index, follow")),
		Meta(Attr("http-equiv", "text/html; charset=utf-8")),
		Meta(Name("msapplication-TileColor"), Content(config.COLOR)),
		Meta(Name("theme-color"), Content(config.COLOR)),
		Meta(Attr("property", "og:url"), Content(config.URL)),
		Meta(Attr("property", "og:type"), Content("website")),
		Meta(Attr("property", "og:title"), Content(config.Title)),
		Meta(Attr("property", "og:description"), Content(config.Description)),
		Meta(Attr("property", "og:image"), Content(config.ImageURL)),
		Meta(Name("twitter:title"), Content(config.Title)),
		Meta(Name("twitter:description"), Content(config.Description)),
		Meta(Name("twitter:image"), Content(config.ImageURL)),
		StyleEl(Raw(pageCSS)),
	)

	page := Doctype(
		HTML(
			head,
			Body(
				H1(Text("hi")),
			),
		),
	)

	return Group{page}.String(), nil
}
