package render

import (
	"context"
	"errors"
	"io"

	"github.com/a-h/templ"
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

		_, err := io.WriteString(w, `<script>`+js+`</script>`)
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
