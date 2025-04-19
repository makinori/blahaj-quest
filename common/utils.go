package common

import (
	"context"
	"encoding/binary"
	"runtime"
)

func Prepend[T any](n T, rest []T) []T {
	return append([]T{n}, rest...)
}

func UniqueHashPC(callerSkip int) string {
	pc, _, _, ok := runtime.Caller(callerSkip)
	if !ok {
		panic("failed to get unique hash")
	}

	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(pc))

	return HashBytes(data)
}

func ChainContextValues(ctx context.Context, values map[any]any) context.Context {
	for k, v := range values {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}

func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}
