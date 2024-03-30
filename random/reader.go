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

var gMReader = MReaderWithRand(Default(), true)

func MReader() io.Reader {
	return gMReader
}

func MReaderWithRand(r *mrand.Rand, safe bool) io.Reader {
	return &mreader{r: r, safe: safe}
}

type mreader struct {
	safe    bool
	readVal uint64
	readPos uint8
	r       *mrand.Rand
}

func (mr *mreader) Read(p []byte) (n int, err error) {
	if mr.safe {
		return ReadBytes(p, mr.r)
	}
	return ReadBytesWithPos(p, mr.r, &mr.readVal, &mr.readPos)
}
