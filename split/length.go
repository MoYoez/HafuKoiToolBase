package split

import "unicode/utf8"

/*

	获取绘制中字符串长度
	获取绘制中字符串长度，并根据字符宽度进行调整。

*/

// GetStringLength
func GetStringLength(words string) float64 {
	var UserFloatNum float64
	var charCount = 0.0
	var truncated string
	for _, runeValue := range words {
		charWidth := utf8.RuneLen(runeValue)
		if charWidth != 3 {
			UserFloatNum = 1.5
		} else {
			UserFloatNum = float64(charWidth)
		}
		truncated += string(runeValue)
		charCount += UserFloatNum
	}
	return charCount
}
