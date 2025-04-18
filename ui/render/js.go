package render

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func JSEl(jsMap map[string]string) Node {
	if len(jsMap) == 0 {
		return nil
	}

	var js string

	for _, snippet := range jsMap {
		js += snippet + "\n"
	}

	return Script(Raw(js))
}
