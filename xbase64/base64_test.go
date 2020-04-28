package xbase64

import (
	"bytes"
	"testing"
)

func Test_Base64(t *testing.T) {
	var cipher bytes.Buffer

	Base64Std.Encode(&cipher, []byte("hello"))

	t.Log(cipher.String())

	var plaintext bytes.Buffer

	Base64Std.Decode(&plaintext, cipher.Bytes())

	t.Log(plaintext.String())
}
