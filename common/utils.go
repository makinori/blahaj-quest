package common

import (
	"context"
	"encoding/binary"
	"hash/crc32"
	"runtime"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Classes(v []string) Node {
	return Class(strings.Join(v, " "))
}

func Attrs(attrs map[string]string) []Node {
	var nodes []Node

	for k, v := range attrs {
		nodes = append(nodes, Attr(k, v))
	}

	return nodes
}

func Prepend[T any](n T, rest []T) []T {
	return append([]T{n}, rest...)
}

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

func UniqueHashPC() string {
	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get unique hash")
	}

	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(pc))

	return HashBytes(data)
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

func ChainContextValues(ctx context.Context, values map[any]any) context.Context {
	for k, v := range values {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}
