package random

import "testing"

func Test_Random(t *testing.T) {
	t.Log(Std.String(10))
	t.Log(Human.String(10))
	t.Log(Text.String(10))
}

func Test_RandomAppendBytes(t *testing.T) {
	var b []byte = []byte("A000_")
	t.Log(string(Std.AppendBytes(b, 9)))
	t.Log(string(Std.AppendBytes(b, 8)))
	t.Log(string(Std.AppendBytes(b, 7)))
	t.Log(string(Std.AppendBytes(b, 6)))
	t.Log(string(Std.AppendBytes(b, 1)))
	t.Log(string(Std.AppendBytes(b, 0)))
}
