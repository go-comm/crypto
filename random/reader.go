package random

import (
	crand "crypto/rand"
	"io"
	mrand "math/rand"
)

var _ = crand.Read
var _ = mrand.New

func CReader() io.Reader {
	return crand.Reader
}

var gMReader = MReaderWithRand(Default())

func MReader() io.Reader {
	return gMReader
}

func MReaderWithRand(r *mrand.Rand) io.Reader {
	return &mreader{r: r}
}

type mreader struct {
	readVal uint64
	readPos uint8
	r       *mrand.Rand
}

func (mr *mreader) Read(p []byte) (n int, err error) {
	return ReadBytesWithPos(p, mr.r, &mr.readVal, &mr.readPos)
}
