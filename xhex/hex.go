package xhex

import (
	"bytes"
	"io"

	"github.com/go-comm/crypto/internal"
)

var (
	HexStd   = NewHexStream(hexStdEncoding)
	HexUpper = NewHexStream(hexUpperEncoding)
)

func NewHexStream(enc Encoding) *HexStream {
	return &HexStream{enc}
}

type HexStream struct {
	Encoding Encoding
}

func (stream *HexStream) Encoder(w io.Writer) io.Writer {
	return NewEncoder(stream.Encoding, w)
}

func (stream *HexStream) WriteStream(w io.Writer, r io.Reader) (err error) {
	ew := NewEncoder(stream.Encoding, w)
	_, err = internal.ZeroCopy(ew, r)
	if err == io.EOF {
		err = nil
	}
	return
}

func (stream *HexStream) Write(w io.Writer, data []byte) (err error) {
	if b, ok := w.(interface{ Grow(n int) }); ok {
		b.Grow(stream.Encoding.EncodedLen(len(data)))
	}
	ew := NewEncoder(stream.Encoding, w)
	ew.Write(data)
	return
}

func (stream *HexStream) WriteToBytes(b []byte, data []byte) []byte {
	buf := bytes.NewBuffer(b)
	stream.Write(buf, data)
	return buf.Bytes()
}

func (stream *HexStream) EncodeToBytes(data []byte) []byte {
	return encode(stream.Encoding, nil, data)
}

func (stream *HexStream) WriteToString(b []byte, data []byte) string {
	return string(stream.WriteToBytes(b, data))
}

func (stream *HexStream) EncodeToString(data []byte) string {
	return string(stream.EncodeToBytes(data))
}

func (stream *HexStream) Decoder(r io.Reader) io.Reader {
	return NewDecoder(stream.Encoding, r)
}

func (stream *HexStream) ReadStream(w io.Writer, r io.Reader) (err error) {
	er := NewDecoder(stream.Encoding, r)
	_, err = internal.ZeroCopy(w, er)
	if err == io.EOF {
		err = nil
	}
	return
}

func (stream *HexStream) Read(b []byte, data io.Reader) ([]byte, error) {
	if l, ok := data.(interface{ Len() int }); ok {
		b = internal.GrowBytes(b, stream.Encoding.DecodedLen(l.Len()))
	}
	er := NewDecoder(stream.Encoding, data)
	return internal.ReadBytes(er, b)
}

func (stream *HexStream) ReadBytes(b []byte, data []byte) ([]byte, error) {
	return stream.Read(b, bytes.NewBuffer(data))
}

func (stream *HexStream) DecodeBytes(data []byte) ([]byte, error) {
	return decode(stream.Encoding, nil, data)
}

func (stream *HexStream) ReadString(b []byte, data string) ([]byte, error) {
	return stream.ReadBytes(b, internal.StringToBytes(&data))
}

func (stream *HexStream) DecodeString(data string) ([]byte, error) {
	return decode(stream.Encoding, nil, internal.StringToBytes(&data))
}
