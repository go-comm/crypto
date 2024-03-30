package xhex

import (
	"encoding/hex"
	"io"

	"github.com/go-comm/crypto/internal"
)

const stdHextable = "0123456789abcdef"
const upperHextable = "0123456789ABCDEF"

var (
	hexStdEncoding   = newEncoding(stdHextable)
	hexUpperEncoding = newEncoding(upperHextable)
)

type Encoding interface {
	Decode(dst []byte, src []byte) (n int, err error)
	DecodedLen(n int) int
	Encode(dst []byte, src []byte) int
	EncodedLen(n int) int
}

func newEncoding(table string) Encoding {
	return &encoding{table: []byte(table)}
}

type encoding struct {
	table []byte
}

func (en *encoding) DecodedLen(n int) int {
	return n / 2
}

func (en *encoding) Decode(dst []byte, src []byte) (n int, err error) {
	return hex.Decode(dst, src)
}

func (en *encoding) EncodedLen(n int) int {
	return n * 2
}

func (en *encoding) Encode(dst []byte, src []byte) int {
	table := en.table
	j := 0
	for _, v := range src {
		dst[j] = table[v>>4]
		dst[j+1] = table[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

func encode(en Encoding, dst []byte, src []byte) []byte {
	encodedLen := en.EncodedLen(len(src))
	if len(dst) < encodedLen {
		dst = make([]byte, encodedLen)
	}
	en.Encode(dst, src)
	return dst[:encodedLen]
}

func decode(en Encoding, dst []byte, src []byte) ([]byte, error) {
	decodedLen := en.DecodedLen(len(src))
	if len(dst) < decodedLen {
		dst = make([]byte, decodedLen)
	}
	n, err := en.Decode(dst, src)
	return dst[:n], err
}

type encoder struct {
	en  Encoding
	w   io.Writer
	err error
}

func NewEncoder(en Encoding, w io.Writer) io.Writer {
	return &encoder{en: en, w: w}
}

func (e *encoder) Write(p []byte) (n int, err error) {
	buf := internal.RequireBuffer()
	defer internal.ReleaseBuffer(buf)

	for len(p) > 0 && e.err == nil {

		chunkSize := len(buf) / 2
		if len(p) < chunkSize {
			chunkSize = len(p)
		}

		var written int
		b := encode(e.en, buf[:], p[:chunkSize])
		written, e.err = e.w.Write(b)
		n += written / 2
		p = p[chunkSize:]
	}
	return n, e.err
}

func NewDecoder(en Encoding, r io.Reader) io.Reader {
	return hex.NewDecoder(r)
}
