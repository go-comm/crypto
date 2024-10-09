package hcode

import (
	"math"

	"github.com/go-comm/crypto/internal"
)

func Int(v int64) uint64 {
	return internal.SumI(internal.Offset, uint64(v))
}

func Uint(v uint64) uint64 {
	return internal.SumI(internal.Offset, v)
}

func Float32(v float32) uint64 {
	return internal.SumI(internal.Offset, uint64(math.Float32bits(v)))
}

func Float64(v float64) uint64 {
	return internal.SumI(internal.Offset, uint64(math.Float64bits(v)))
}

func String(s string) uint64 {
	return internal.SumB(internal.Offset, []byte(s))
}

func Bytes(b []byte) uint64 {
	return internal.SumB(internal.Offset, b)
}
