package xhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

func Hash(h hash.Hash, b []byte) string {
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func Sum(h hash.Hash, b []byte, d []byte) []byte {
	h.Write(b)
	return h.Sum(d)
}

func Sum32(h hash.Hash32, b []byte) uint32 {
	h.Write(b)
	return h.Sum32()
}

func Sum64(h hash.Hash64, b []byte) uint64 {
	h.Write(b)
	return h.Sum64()
}

func Md5(text string) string {
	return Hash(md5.New(), []byte(text))
}

func Sha1(text string) string {
	return Hash(sha1.New(), []byte(text))
}

func Sha256(text string) string {
	return Hash(sha256.New(), []byte(text))
}

func Sha512(text string) string {
	return Hash(sha512.New(), []byte(text))
}

func Crc32IEEE(text string) uint32 {
	return Sum32(crc32.NewIEEE(), []byte(text))
}

func Crc64ISO(text string) uint64 {
	return Sum64(crc64.New(crc64.MakeTable(crc64.ISO)), []byte(text))
}

func Crc64ECMA(text string) uint64 {
	return Sum64(crc64.New(crc64.MakeTable(crc64.ECMA)), []byte(text))
}

func Adler32(text string) uint32 {
	return Sum32(adler32.New(), []byte(text))
}

func Fnv32(text string) uint32 {
	return Sum32(fnv.New32(), []byte(text))
}

func Fnv32a(text string) uint32 {
	return Sum32(fnv.New32a(), []byte(text))
}

func Fnv64(text string) uint64 {
	return Sum64(fnv.New64(), []byte(text))
}

func Fnv64a(text string) uint64 {
	return Sum64(fnv.New64a(), []byte(text))
}

func AppendUint32(b []byte, x uint32) []byte {
	a := [4]byte{
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
	return append(b, a[:]...)
}

func ReadUint32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func AppendUint64(b []byte, x uint64) []byte {
	a := [8]byte{
		byte(x >> 56),
		byte(x >> 48),
		byte(x >> 40),
		byte(x >> 32),
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
	return append(b, a[:]...)
}

func ReadUint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}
