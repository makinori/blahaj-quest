package render

import (
	"bytes"
	"context"
	"errors"

	"github.com/makinori/blahaj-quest/common"

	"github.com/a-h/templ"
)

func Render(
	ctx context.Context,
	layout func(body string) templ.Component,
	page func() templ.Component,
) ([]byte, error) {
	ctx = common.ChainContextValues(
		ctx,
		map[any]any{
			pageSCSSKey:   make(map[string]string),
			pageHeadJSKey: make(map[string]string),
			pageBodyJSKey: make(map[string]string),
		},
	)

	buf := bytes.NewBuffer(nil)

	err := page().Render(ctx, buf)
	if err != nil {
		return []byte{}, errors.New("failed to render page: " + err.Error())
	}

	pageHTML := buf.String()

	buf = bytes.NewBuffer(nil)

	err = layout(pageHTML).Render(ctx, buf)
	if err != nil {
		return []byte{}, errors.New("failed to render layout: " + err.Error())
	}

	// return buf.Bytes(), nil

	html, err := common.MinifyHTML(buf.Bytes())
	if err != nil {
		return []byte{}, errors.New("failed to minify: " + err.Error())
	}

	return html, nil
}
