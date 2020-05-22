package xbase64

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/go-comm/xbytes"
)

var (
	Base64Std = NewBase64Stream(base64.StdEncoding)
	Base64URL = NewBase64Stream(base64.URLEncoding)
)

func NewBase64Stream(enc *base64.Encoding) *Base64Stream {
	return &Base64Stream{enc}
}

type Base64Stream struct {
	Encoding *base64.Encoding
}

func (e *Base64Stream) EncodeFromReader(w io.Writer, r io.Reader) error {
	ew := base64.NewEncoder(e.Encoding, w)
	xbytes.ZeroCopy(ew, r)
	return ew.Close()
}

func (e *Base64Stream) Encode(w io.Writer, data []byte) error {
	if buf, ok := w.(*bytes.Buffer); ok {
		buf.Grow(e.Encoding.EncodedLen(len(data)))
	}
	ew := base64.NewEncoder(e.Encoding, w)
	ew.Write(data)
	return ew.Close()
}

func (e *Base64Stream) Marshal(dst []byte, src []byte) ([]byte, error) {
	encodedLen := e.Encoding.EncodedLen(len(src))
	if len(dst) < encodedLen {
		dst = make([]byte, encodedLen)
	}
	e.Encoding.Encode(dst, src)
	return dst[:encodedLen], nil
}

func (e *Base64Stream) DecodeFromReader(w io.Writer, r io.Reader) error {
	er := base64.NewDecoder(e.Encoding, r)
	_, err := xbytes.ZeroCopy(w, er)
	return err
}

func (e *Base64Stream) Decode(w io.Writer, data []byte) error {
	if buf, ok := w.(*bytes.Buffer); ok {
		buf.Grow(e.Encoding.DecodedLen(len(data)))
	}
	return e.DecodeFromReader(w, bytes.NewReader(data))
}

func (e *Base64Stream) DecodeToBytes(s *string) ([]byte, error) {
	buf := xbytes.StringToBytes(s)
	dbuf := make([]byte, e.Encoding.DecodedLen(len(buf)))
	n, err := e.Encoding.Decode(dbuf, buf)
	return dbuf[:n], err
}

func (e *Base64Stream) Unmarshal(dst []byte, src []byte) ([]byte, error) {
	decodedLen := e.Encoding.DecodedLen(len(src))
	if len(dst) < decodedLen {
		dst = make([]byte, decodedLen)
	}
	n, err := e.Encoding.Decode(dst, src)
	return dst[:n], err
}
