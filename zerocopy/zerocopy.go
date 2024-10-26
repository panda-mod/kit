package zerocopy

import "unsafe"

// String 零拷贝将字节切片转换为字符串ß
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes 零拷贝将字符串转换为字节切片
func Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
