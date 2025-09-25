package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var hexColorRegexp = regexp.MustCompile(
	`(?i)^#?([a-f0-9]{1,2})([a-f0-9]{1,2})([a-f0-9]{1,2})`,
)

type Color struct {
	R uint8
	G uint8
	B uint8
	// A uint8
}

func ParseHexColor(hexColor string) (Color, error) {
	// doesnt handle #000

	matches := hexColorRegexp.FindStringSubmatch(hexColor)
	if len(matches) == 0 {
		return Color{}, errors.New("failed to parse hex color")
	}

	var color Color

	channels := []*uint8{
		&color.R,
		&color.G,
		&color.B,
	}

	for i, channel := range channels {
		i += 1

		if len(matches[i]) == 1 {
			matches[i] += matches[1]
		}

		value, _ := strconv.ParseUint(matches[i], 16, 8)
		*channel = uint8(value)
	}

	return color, nil
}

func ColorToHex(color Color) string {
	return fmt.Sprintf("#%02x%02x%02x", color.R, color.G, color.B)
}

func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

func MixHexColors(aHex string, bHex string, t float64) string {
	a, err := ParseHexColor(aHex)
	if err != nil {
		return ""
	}

	b, err := ParseHexColor(bHex)
	if err != nil {
		return ""
	}

	return ColorToHex(Color{
		R: uint8(Lerp(float64(a.R), float64(b.R), t)),
		G: uint8(Lerp(float64(a.G), float64(b.G), t)),
		B: uint8(Lerp(float64(a.B), float64(b.B), t)),
	})
}
