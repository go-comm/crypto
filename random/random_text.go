package random

import (
	crand "crypto/rand"
	"io"
	mrand "math/rand"

	"github.com/go-comm/crypto/internal"
)

var _ = crand.Read
var _ = mrand.New

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
	Number      = NewRandomText(numberTable)
	Upper       = NewRandomText(upperTable)
	NumberUpper = NewRandomText(MergeTable(numberTable, upperTable))
	Lower       = NewRandomText(lowerTable)
	NumberLower = NewRandomText(MergeTable(numberTable, lowerTable))
	Alpha       = NewRandomText(alphaTable)
	Std         = NewRandomText(stdTable)
	Human       = NewRandomText(humanTable)
	Text        = NewRandomText(textTable)
	R16         = NewRandomText(textTable[:16])
	R32         = NewRandomText(textTable[:32])
	R48         = NewRandomText(textTable[:48])
)

func MergeTable(a []byte, b []byte, c ...[]byte) []byte {
	var d []byte
	d = append(d, a...)
	d = append(d, b...)
	if len(c) > 0 {
		for _, e := range c {
			d = append(d, e...)
		}
	}
	return d
}

func NewRandomText(table []byte) *RandomText {
	if len(table) >= 256 {
		panic("random: The length of table must be less than 256")
	}
	rt := &RandomText{table: table, rr: MReader()}
	return rt
}

type RandomText struct {
	table []byte
	rr    io.Reader
}

func (rt RandomText) WithoutCrypto() *RandomText {
	return rt.SetReader(MReader())
}

func (rt RandomText) WithCrypto() *RandomText {
	return rt.SetReader(CReader())
}

func (rt RandomText) SetReader(rr io.Reader) *RandomText {
	rt.rr = rr
	return &rt
}

func (rt *RandomText) read(rr io.Reader, b []byte) (n int, err error) {
	_, err = io.ReadFull(rr, b)
	if err != nil {
		return 0, err
	}
	size := len(rt.table)
	n = len(b)
	for i := 0; i < n; i++ {
		b[i] = rt.table[int(b[i])*size>>8]
	}
	return
}

func (rt *RandomText) Read(b []byte) (n int, err error) {
	n, err = rt.read(rt.rr, b)
	if err == nil {
		return
	}
	n, err = rt.read(MReader(), b)
	if err == nil {
		return
	}
	panic(err)
}

func (rt *RandomText) Bytes(n int) []byte {
	b := make([]byte, n)
	rt.Read(b)
	return b
}

func (rt *RandomText) AppendBytes(b []byte, n int) []byte {
	b = internal.GrowBytes(b, n)
	var tmp [8]byte
	i := n - 1
	for ; i >= 7; i -= 8 {
		rt.Read(tmp[:])
		b = append(b, tmp[:]...)
	}
	if i >= 0 {
		rt.Read(tmp[:])
		b = append(b, tmp[:i+1]...)
	}
	return b
}

func (rt *RandomText) String(n int) string {
	return string(rt.Bytes(n))
}

func (rt *RandomText) RandomToString(n int) string {
	return string(rt.Bytes(n))
}
