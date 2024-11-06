package b2s

// COPY FROM https://github.com/wdvxdr1123/ZeroBot/blob/main/utils/helper/helper.go

import "unsafe"

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 没有内存开销的转换
func StringToBytes(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
