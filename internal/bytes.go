package internal

import (
	"io"
	"sync"
	"unsafe"
)

func GrowBytes(b []byte, n int) []byte {
	if cap(b)-len(b) < n {
		d := make([]byte, 0, len(b)+n)
		d = append(d, b...)
		return d
	}
	return b
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StringToBytes(s *string) []byte {
	bs := (*[2]uintptr)(unsafe.Pointer(s))
	b := [3]uintptr{bs[0], bs[1], bs[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

var buffer = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 512)
	},
}

func RequireBuffer() []byte {
	return buffer.Get().([]byte)
}

func ReleaseBuffer(buf []byte) {
	buffer.Put(buf)
}

func ZeroCopy(dst io.Writer, src io.Reader) (int64, error) {
	buf := RequireBuffer()
	defer ReleaseBuffer(buf)
	return io.CopyBuffer(dst, src, buf)
}

func ReadBytes(r io.Reader, b []byte) ([]byte, error) {
	if len(b) == 0 {
		b = make([]byte, 0, 512)
	}
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
	}
}
