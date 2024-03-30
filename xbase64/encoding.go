package xbase64

type Encoding interface {
	Decode(dst []byte, src []byte) (n int, err error)
	DecodeString(s string) ([]byte, error)
	DecodedLen(n int) int
	Encode(dst []byte, src []byte)
	EncodeToString(src []byte) string
	EncodedLen(n int) int
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
