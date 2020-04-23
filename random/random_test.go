package random

import "testing"

func Test_Random(t *testing.T) {
	t.Log(Std.RandomToString(10))
	t.Log(Human.RandomToString(10))
	t.Log(Text.RandomToString(10))
}
