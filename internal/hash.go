package internal

const (
	Offset uint64 = 14695981039346656037
	Prime  uint64 = 1099511628211
)

func SumI(h uint64, v uint64) uint64 {
	h *= Prime
	h ^= v
	return h
}

func SumB(h uint64, s []byte) uint64 {
	for i := 0; i < len(s); i++ {
		h *= Prime
		h ^= uint64(s[i])
	}
	return h
}
