package random

import "testing"

func Test_Random(t *testing.T) {
	t.Log(Std.RandomToString(10))
	t.Log(Human.RandomToString(10))
	t.Log(Text.RandomToString(10))
}

func Test_RandomWithCrypto(t *testing.T) {
	t.Log(Std.RandomToString(10))
	t.Log(Std.WithCrypto().RandomToString(10))
	t.Log(Std.RandomToString(10))

	t.Log(Std == Std)
	t.Log(Std != Std.WithCrypto())
}
