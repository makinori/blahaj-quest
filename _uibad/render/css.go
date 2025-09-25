package render

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/makinori/blahaj-quest/common"

	"github.com/a-h/templ"
	sass "github.com/bep/godartsass/v2"
	"github.com/charmbracelet/log"
)

var (
	pageSCSSKey = "pageSCSS"

	sassTranspiler *sass.Transpiler
)

func ensureSass() {
	if sassTranspiler != nil {
		return
	}

	var err error
	sassTranspiler, err = sass.Start(sass.Options{})

	if err != nil {
		panic(err)
	}
}

func SCSS(ctx context.Context, input string) string {
	pageSCSS, ok := ctx.Value(pageSCSSKey).(map[string]string)

	if !ok {
		log.Error("failed to get page scss from context")
		return ""
	}

	// var source string

	// // run input through sass first?

	// for line := range strings.SplitSeq(input, "\n") {
	// 	line = strings.TrimSpace(line)
	// 	if line == "" {
	// 		continue
	// 	}
	// 	source += line + "\n"
	// }

	className := common.HashString(input)

	pageSCSS[className] = input

	return className
}

func SCSSElement(extraScss ...string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		pageSCSS, ok := ctx.Value(pageSCSSKey).(map[string]string)
		if !ok {
			return errors.New("failed to get page scss from context")
		}

		var source string

		for _, snippet := range extraScss {
			source += snippet + "\n"
		}

		for className, snippet := range pageSCSS {
			source += "." + className + "{" + snippet + "} "
		}

		source = strings.TrimSpace(source)

		ensureSass()

		res, err := sassTranspiler.Execute(sass.Args{
			Source:          source,
			OutputStyle:     sass.OutputStyleCompressed,
			SourceSyntax:    sass.SourceSyntaxSCSS,
			EnableSourceMap: false,
		})

		if err != nil {
			return errors.New("failed to compile scss")
		}

		_, err = io.WriteString(w, `<style type="text/css">`+res.CSS+`</style>`)
		if err != nil {
			return err
		}

		return nil
	})
}
