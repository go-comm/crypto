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

func NewBase64Stream(enc *base64.Encoding) *base64Stream {
	return &base64Stream{enc: enc}
}

type base64Stream struct {
	enc *base64.Encoding
}

func (e *base64Stream) EncodeFromReader(w io.Writer, r io.Reader) error {
	ew := base64.NewEncoder(e.enc, w)
	xbytes.ZeroCopy(ew, r)
	return ew.Close()
}

func (e *base64Stream) Encode(w io.Writer, data []byte) error {
	if buf, ok := w.(*bytes.Buffer); ok {
		buf.Grow(e.enc.EncodedLen(len(data)))
	}
	ew := base64.NewEncoder(e.enc, w)
	ew.Write(data)
	return ew.Close()
}

func (e *base64Stream) DecodeFromReader(w io.Writer, r io.Reader) error {
	er := base64.NewDecoder(e.enc, r)
	_, err := xbytes.ZeroCopy(w, er)
	return err
}

func (e *base64Stream) Decode(w io.Writer, data []byte) error {
	if buf, ok := w.(*bytes.Buffer); ok {
		buf.Grow(e.enc.DecodedLen(len(data)))
	}
	return e.DecodeFromReader(w, bytes.NewReader(data))
}

func (e *base64Stream) DecodeToBytes(s *string) ([]byte, error) {
	buf := xbytes.StringToBytes(s)
	dbuf := make([]byte, e.enc.DecodedLen(len(buf)))
	n, err := e.enc.Decode(dbuf, buf)
	return dbuf[:n], err
}
