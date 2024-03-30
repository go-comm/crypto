package xbase64

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/go-comm/crypto/internal"
)

var (
	Base64Std    = NewBase64Stream(base64.StdEncoding)
	Base64RawStd = NewBase64Stream(base64.StdEncoding.WithPadding(base64.NoPadding))
	Base64URL    = NewBase64Stream(base64.URLEncoding)
	Base64RawURL = NewBase64Stream(base64.URLEncoding.WithPadding(base64.NoPadding))
)

func NewBase64Stream(enc *base64.Encoding) *Base64Stream {
	return &Base64Stream{enc}
}

type Base64Stream struct {
	Encoding *base64.Encoding
}

func (stream *Base64Stream) Encoder(w io.WriteCloser) io.WriteCloser {
	return base64.NewEncoder(stream.Encoding, w)
}

func (stream *Base64Stream) WriteStream(w io.Writer, r io.Reader) (err error) {
	ew := base64.NewEncoder(stream.Encoding, w)
	internal.ZeroCopy(ew, r)
	err = ew.Close()
	if err == io.EOF {
		err = nil
	}
	return
}

func (stream *Base64Stream) Write(w io.Writer, data []byte) (err error) {
	if b, ok := w.(interface{ Grow(n int) }); ok {
		b.Grow(stream.Encoding.EncodedLen(len(data)))
	}
	ew := base64.NewEncoder(stream.Encoding, w)
	ew.Write(data)
	err = ew.Close()
	return
}

func (stream *Base64Stream) WriteToBytes(b []byte, data []byte) []byte {
	buf := bytes.NewBuffer(b)
	stream.Write(buf, data)
	return buf.Bytes()
}

func (stream *Base64Stream) EncodeToBytes(data []byte) []byte {
	return encode(stream.Encoding, nil, data)
}

func (stream *Base64Stream) WriteToString(b []byte, data []byte) string {
	return string(stream.WriteToBytes(b, data))
}

func (stream *Base64Stream) EncodeToString(data []byte) string {
	return string(encode(stream.Encoding, nil, data))
}

func (stream *Base64Stream) Decoder(r io.Reader) io.Reader {
	return base64.NewDecoder(stream.Encoding, r)
}

func (stream *Base64Stream) ReadStream(w io.Writer, r io.Reader) (err error) {
	er := base64.NewDecoder(stream.Encoding, r)
	_, err = internal.ZeroCopy(w, er)
	if err == io.EOF {
		err = nil
	}
	return
}

func (stream *Base64Stream) Read(b []byte, data io.Reader) ([]byte, error) {
	if l, ok := data.(interface{ Len() int }); ok {
		b = internal.GrowBytes(b, stream.Encoding.DecodedLen(l.Len()))
	}
	er := base64.NewDecoder(stream.Encoding, data)
	return internal.ReadBytes(er, b)
}

func (stream *Base64Stream) ReadBytes(b []byte, data []byte) ([]byte, error) {
	return stream.Read(b, bytes.NewBuffer(data))
}

func (stream *Base64Stream) DecodeBytes(data []byte) ([]byte, error) {
	return decode(stream.Encoding, nil, data)
}

func (stream *Base64Stream) ReadString(b []byte, data string) ([]byte, error) {
	return stream.ReadBytes(b, internal.StringToBytes(&data))
}

func (stream *Base64Stream) DecodeString(data string) ([]byte, error) {
	return decode(stream.Encoding, nil, internal.StringToBytes(&data))
}
