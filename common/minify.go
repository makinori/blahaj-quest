package common

import (
	"github.com/dtrenin7/minify/v2"
	"github.com/dtrenin7/minify/v2/html"
	"github.com/dtrenin7/minify/v2/js"
)

var minifier *minify.M

func ensureMinifier() {
	if minifier != nil {
		return
	}

	minifier = minify.New()

	minifier.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
	})

	minifier.Add("application/javascript", &js.Minifier{})
}

func MinifyHTML(input []byte) ([]byte, error) {
	ensureMinifier()

	html, err := minifier.Bytes("text/html", input)
	if err != nil {
		return []byte{}, err
	}

	return html, nil
}

func MinifyJS(input string) (string, error) {
	ensureMinifier()

	js, err := minifier.String("application/javascript", input)
	if err != nil {
		return "", err
	}

	return js, nil
}
