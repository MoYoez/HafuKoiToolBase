package split

import "unicode/utf8"

/*

	分割中文字符串，按照中文字符长度分割。
	这个是为了避免中文字符串过长，在一部份地方渲染的时候，因为长度会导致一部分字符串无法显示（错位）。
	所以说需要指定一个长度，然后按照这个长度进行分割，这样好方便渲染。

*/

// SplitChineseString
func SplitChineseString(s string, length int) []string {
	results := make([]string, 0)
	runes := []rune(s)
	start := 0
	for i := 0; i < len(runes); i++ {
		size := utf8.RuneLen(runes[i])
		if start+size > length {
			results = append(results, string(runes[0:i]))
			runes = runes[i:]
			i, start = 0, 0
		}
		start += size
	}
	if len(runes) > 0 {
		results = append(results, string(runes))
	}
	return results
}
