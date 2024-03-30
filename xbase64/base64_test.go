package xbase64

import (
	"bytes"
	"testing"
)

func Test_Base64Stream(t *testing.T) {
	var cipher bytes.Buffer

	Base64Std.WriteStream(&cipher, bytes.NewBufferString("hello"))

	t.Log(cipher.String())

	var plaintext bytes.Buffer

	Base64Std.ReadStream(&plaintext, &cipher)

	t.Log(plaintext.String())
}

func Test_EncodeToBytes(t *testing.T) {

	var b []byte = []byte("user_")
	b = Base64Std.WriteToBytes(b, []byte("john"))
	t.Log(string(b))
}
