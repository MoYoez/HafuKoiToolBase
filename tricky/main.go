package tricky

import "strings"

/*

	字符串处理
	将字符串转换为切片，或将字符串转换为布尔值。

*/

func MakeStrintoSlice(str string) []string {
	return strings.Split(str, ",")
}

func MakeStrIntoBool(str string) bool {
	str = strings.ToLower(strings.TrimSpace(str))
	return str == "true" || str == "1" || str == "yes" || str == "y"
}
