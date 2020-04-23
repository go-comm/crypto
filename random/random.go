package random

import (
	"crypto/rand"
	"fmt"
	"unsafe"
)

// slice's length must be greater than 0
func checkLength(length int) {
	if length <= 0 {
		panic(fmt.Sprintf("[checkLength] expect length > 0, but not %v", length))
	}
}

// generate byte slice
func GenerateBytes(length int) []byte {
	checkLength(length)
	b := make([]byte, length)
	rand.Read(b)
	return b
}

var (
	stdTable   = []byte("0123456789ABCDEF")
	humanTable = []byte("346789ABCDEFGHIJKLMNPQRTUVWXYabcdefghijkmnpqrtuvwxy")
	textTable  = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

var (
	Std   = NewRandom(stdTable)
	Human = NewRandom(humanTable)
	Text  = NewRandom(textTable)
)

type Random struct {
	table []byte
}

func NewRandom(table []byte) *Random {
	r := &Random{table: table}
	return r
}

func (r *Random) RandomToString(n int) string {
	buf := GenerateBytes(n)
	tableLen := len(r.table)
	for i := 0; i < n; i++ {
		buf[i] = r.table[int(buf[i])*tableLen>>8]
	}
	return *(*string)(unsafe.Pointer(&buf))
}
