package random

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"os"
	"time"
	"unsafe"
)

var rnd = mrand.New(mrand.NewSource(time.Now().UnixNano() + int64(os.Getpid())<<10))

// slice's length must be greater than 0
func checkLength(length int) {
	if length <= 0 {
		panic(fmt.Sprintf("[checkLength] expect length > 0, but not %v", length))
	}
}

// generate byte slice
func GenerateCryptoBytes(length int) []byte {
	checkLength(length)
	b := make([]byte, length)
	crand.Read(b)
	return b
}

// generate byte slice
func GenerateBytes(length int) []byte {
	checkLength(length)
	b := make([]byte, length)
	var i int
	var v int64
	for i = length - 1; i >= 7; i -= 8 {
		v = rnd.Int63()
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
		b[i-3] = byte(v >> 24)
		b[i-4] = byte(v >> 32)
		b[i-5] = byte(v >> 40)
		b[i-6] = byte(v >> 48)
		b[i-7] = byte(v >> 56)
	}
	if i < 0 {
		return b
	}
	v = rnd.Int63()
	switch i {
	case 6:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
		b[i-3] = byte(v >> 24)
		b[i-4] = byte(v >> 32)
		b[i-5] = byte(v >> 40)
		b[i-6] = byte(v >> 48)
	case 5:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
		b[i-3] = byte(v >> 24)
		b[i-4] = byte(v >> 32)
		b[i-5] = byte(v >> 40)
	case 4:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
		b[i-3] = byte(v >> 24)
		b[i-4] = byte(v >> 32)
	case 3:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
		b[i-3] = byte(v >> 24)
	case 2:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
		b[i-2] = byte(v >> 16)
	case 1:
		b[i] = byte(v)
		b[i-1] = byte(v >> 8)
	case 0:
		b[i] = byte(v)
	}
	return b
}

var (
	numberTable = []byte("0123456789")
	upperTable  = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerTable  = []byte("abcdefghijklmnopqrstuvwxyz")
	alphaTable  = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	textTable   = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	stdTable    = []byte("0123456789ABCDEF")
	humanTable  = []byte("346789ABCDEFGHIJKLMNPQRTUVWXYabcdefghijkmnpqrtuvwxy")
)

var (
	Number = NewRandom(numberTable)
	Upper  = NewRandom(upperTable)
	lower  = NewRandom(lowerTable)
	Alpha  = NewRandom(alphaTable)
	Std    = NewRandom(stdTable)
	Human  = NewRandom(humanTable)
	Text   = NewRandom(textTable)
)

func NewRandom(table []byte) *Random {
	if len(table) >= 256 {
		panic("random: The length of table must be less than 256")
	}
	r := &Random{table: table, genbytes: GenerateBytes}
	return r
}

type Random struct {
	table    []byte
	genbytes func(length int) []byte
}

func (r Random) WithCrypto() *Random {
	r.genbytes = GenerateCryptoBytes
	return &r
}

func (r *Random) RandomToString(n int) string {
	buf := r.genbytes(n)
	tableLen := len(r.table)
	for i := 0; i < n; i++ {
		buf[i] = r.table[int(buf[i])*tableLen>>8]
	}
	return *(*string)(unsafe.Pointer(&buf))
}
