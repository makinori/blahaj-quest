package render

import (
	"context"
	"errors"
	"io"

	"github.com/a-h/templ"
	"github.com/makinori/blahaj-quest/common"
)

var (
	pageHeadJSKey = "pageHeadJS"
	pageBodyJSKey = "pageBodyJS"
)

func jsElement(kind string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		pageJS, ok := ctx.Value(kind).(map[string]string)
		if !ok {
			return errors.New("failed to get " + kind + " from context")
		}

		if len(pageJS) == 0 {
			return nil
		}

		var js string

		for _, snippet := range pageJS {
			js += snippet + "\n"
		}

		js, err := common.MinifyJS(js)
		if err != nil {
			return err
		}

		_, err = io.WriteString(w, `<script>`+js+`</script>`)
		if err != nil {
			return err
		}

		return nil
	})
}

func HeadJSElement() templ.Component {
	return jsElement(string(pageHeadJSKey))
}

func BodyJSElement() templ.Component {
	return jsElement(string(pageBodyJSKey))
}

func addJS(ctx context.Context, kind string, id string, js string) error {
	pageJS, ok := ctx.Value(kind).(map[string]string)
	if !ok {
		return errors.New("failed to get " + kind + " from context")
	}

	pageJS[id] = js

	return nil
}

func AddHeadJS(ctx context.Context, id string, js string) error {
	return addJS(ctx, pageHeadJSKey, id, js)
}

func AddBodyJS(ctx context.Context, id string, js string) error {
	return addJS(ctx, pageBodyJSKey, id, js)
}
