package ui

import (
	_ "embed"

	. "github.com/makinori/blahaj-quest/common"
	. "github.com/makinori/blahaj-quest/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed layout.scss
var layoutStyles string

func Layout(r *RenderContext, children ...Node) Node {

	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				TitleEl(Text(ConfigTitle)),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=0.7")),
				Link(Attrs(map[string]string{
					"rel":   "icon",
					"type":  "image/png",
					"sizes": "16x16",
					"href":  "/favicon-16x16.png",
				})...),
				Link(Attrs(map[string]string{
					"rel":   "icon",
					"type":  "image/png",
					"sizes": "32x32",
					"href":  "/favicon-32x32.png",
				})...),
				Meta(Name("title"), Content(ConfigTitle)),
				Meta(Name("description"), Content(ConfigDescription)),
				Meta(Name("keywords"), Content(ConfigKeywords)),
				Meta(Name("robots"), Content("index, follow")),
				Meta(Attrs(map[string]string{
					"http-equiv": "text/html; charset=utf-8",
				})...),
				Meta(Name("msapplication-TileColor"), Content(ConfigColor)),
				Meta(Name("theme-color"), Content(ConfigColor)),
				Meta(Attr("property", "og:url"), Content(ConfigURL)),
				Meta(Attr("property", "og:type"), Content("website")),
				Meta(Attr("property", "og:title"), Content(ConfigTitle)),
				Meta(Attr("property", "og:description"), Content(ConfigDescription)),
				Meta(Attr("property", "og:image"), Content(ConfigImageURL)),
				Meta(Attr("property", "twitter:url"), Content(ConfigURL)),
				Meta(Attr("name", "twitter:title"), Content(ConfigTitle)),
				Meta(Attr("name", "twitter:description"), Content(ConfigDescription)),
				Meta(Attr("name", "twitter:image"), Content(ConfigImageURL)),
				SCSSEl(r, layoutStyles),
				JSEl(r.HeadJS),
			),
			Body(append(
				children,
				JSEl(r.BodyJS),
			)...),
		),
	)
}
