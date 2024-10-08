package random

import (
	crand "crypto/rand"
	"math/bits"
	mrand "math/rand"
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

var _ = crand.Read
var _ = mrand.New

const offset uint64 = 14695981039346656037
const prime uint64 = 1099511628211
const rngMask uint64 = 1<<63 - 1

func seed() uint64 {
	h := offset
	h = hash(h, uint64(time.Now().UnixNano()))
	h = hash(h, uint64(os.Getpid()+1))
	return h
}

func hash(v uint64, salt uint64) uint64 {
	v *= salt
	v ^= prime
	return v
}

func fastrand64(n *uint64) uint64 {
	v := atomic.AddUint64(n, 0xa0761d6478bd642f)
	hi, lo := bits.Mul64(v, v^0xe7037ed1a0b428db)
	return hi ^ lo
}

func NewSeed() int64 {
	return int64(seed() & rngMask)
}

var gRand = NewFastRand()

func Default() *mrand.Rand {
	return gRand
}

func Seed(seed int64) {
	gRand.Seed(seed)
}

func Float32() float32 {
	return gRand.Float32()
}

func Float64() float64 {
	return gRand.Float64()
}

func Int() int {
	return gRand.Int()
}

func Int31() int32 {
	return gRand.Int31()
}

func Int31n(n int32) int32 {
	return gRand.Int31n(n)
}

func Int63() int64 {
	return gRand.Int63()
}

func Int63n(n int64) int64 {
	return gRand.Int63n(n)
}

func Intn(n int) int {
	return gRand.Intn(n)
}

func NormFloat64() float64 {
	return gRand.NormFloat64()
}

func Perm(n int) []int {
	return gRand.Perm(n)
}

func Shuffle(n int, swap func(i, j int)) {
	gRand.Shuffle(n, swap)
}

func Uint32() uint32 {
	return gRand.Uint32()
}

func Uint64() uint64 {
	return gRand.Uint64()
}

func ExpFloat64() float64 {
	return gRand.ExpFloat64()
}

func Read(p []byte) (n int, err error) {
	return ReadBytes(p, gRand)
}

func NewFastRand() *mrand.Rand {
	return mrand.New(NewFastSource())
}

func NewFastRandWithSeed(seed int64) *mrand.Rand {
	return mrand.New(NewFastSourceWithSeed(seed))
}

func NewFastSource() mrand.Source {
	return NewFastSourceWithSeed(NewSeed())
}

func NewFastSourceWithSeed(seed int64) mrand.Source {
	src := &fastSource{}
	src.Seed(seed)
	return src
}

type fastSource struct {
	state [3]uint32
}

func (src *fastSource) r() *uint64 {
	if uintptr(unsafe.Pointer(&src.state))%8 == 0 {
		return (*uint64)(unsafe.Pointer(&src.state[0]))
	}
	return (*uint64)(unsafe.Pointer(&src.state[1]))
}

func (src *fastSource) Seed(seed int64) {
	atomic.StoreUint64(src.r(), uint64(seed&int64(rngMask)))
}

func (src *fastSource) Int63() (n int64) {
	return int64(src.Uint64() & rngMask)
}

func (src *fastSource) Uint64() (n uint64) {
	return fastrand64(src.r())
}

type Source64 interface {
	mrand.Source
	Uint64() uint64
}

func ReadBytesWithPos(p []byte, src Source64, readVal *uint64, readPos *uint8) (n int, err error) {
	val := *readVal
	pos := *readPos
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val = src.Uint64()
			pos = 7
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}
	*readPos = pos
	*readVal = val
	return
}

func ReadBytes(p []byte, src Source64) (n int, err error) {
	var readVal uint64 = 0
	var readPos uint8 = 0
	return ReadBytesWithPos(p, gRand, &readVal, &readPos)
}
