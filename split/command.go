package split

import "strings"

/*

	分割命令
	分割命令， 可根据需要长度分割，如果长度切割过少，则剩下部分采取不截断。

*/

// SplitCommandTo Split Command and Adjust To.
func SplitCommandTo(raw string, setCommandStopper int) (splitCommandLen int, splitInfo []string) {
	rawSplit := strings.SplitN(raw, " ", setCommandStopper)
	return len(rawSplit), rawSplit
}
