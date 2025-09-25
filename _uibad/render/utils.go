package render

import (
	"os"
	"path"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/makinori/blahaj-quest/common"
)

func DevEmbed(variable *string, filename string) {
	if !common.ConfigInDev {
		return
	}

	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		log.Error("failed to get caller")
		return
	}

	filePath = path.Join(path.Dir(filePath), filename)

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Error("failed to read file: " + filename)
		return
	}

	*variable = string(bytes)
}
