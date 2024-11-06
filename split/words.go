package split

import "unicode/utf8"

/*

	断词器
	将字符串断词，根据长度截断省略。

*/

// BreakWords Break words into pieces and make adjust here.
func BreakWords(breakWords string, LimitedLength float64) string {
	var setBreaker bool
	var charCount = 0.0
	var truncated string
	var UserFloatNum float64
	for _, runeValue := range breakWords {
		charWidth := utf8.RuneLen(runeValue)
		if charWidth != 3 {
			UserFloatNum = 1.5
		} else {
			UserFloatNum = float64(charWidth)
		}
		if charCount+UserFloatNum > LimitedLength {
			setBreaker = true
			break
		}
		truncated += string(runeValue)
		charCount += UserFloatNum
	}
	if setBreaker {
		return truncated + "..."
	} else {
		return truncated
	}
}
