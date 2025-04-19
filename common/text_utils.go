package common

import (
	"fmt"
	"hash/crc32"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base52 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func HashBytes(data []byte) string {
	hash := crc32.ChecksumIEEE(data)
	if hash == 0 {
		return string(base52[0])
	}

	var name string
	for hash > 0 {
		name = string(base52[hash%52]) + name
		hash /= 52
	}

	return name
}

func HashString(data string) string {
	return HashBytes([]byte(data))
}

func FormatNumber(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d\n", n)
}

func Plural(n int, single string, plural ...string) string {
	if n == 1 || n == -1 {
		return FormatNumber(n) + " " + single
	}
	if len(plural) > 0 {
		return FormatNumber(n) + " " + plural[0]
	} else {
		return FormatNumber(n) + " " + single + "s"
	}
}

func LastUpdated(lastUpdated time.Time) string {
	t := time.Since(lastUpdated)

	h := t.Hours()
	m := t.Minutes()

	if m < 60 {
		return fmt.Sprintf("%.0f minutes ago", m)
	}

	if h < 24 {
		return fmt.Sprintf("%.0f hours ago", h)
	}

	return fmt.Sprintf("%.0f days ago", h/24)
}
