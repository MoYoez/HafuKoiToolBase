package passkey

/*

	基于凯撒密码和自建加密算法，实现了一套简易的加密解密方案，当然了，他没有双边混淆（ 双边无需指定Hash，可以给Hash价格 Offset 偏移 和 简单的AES 来解决问题），所以说算是个简易的加密解密方案。

	这一套本身是给某家公司写得玩意（其目的是想让一些支付接口仅可用一次，不想滥用，而且因为这个上数据库没必要 ），但是呢，这家公司是毁约了，所以这套东西就流落出来了。

	反正最后是个遗弃函数，实际上也没用上。

	不过，如果需要的话，可以参考一下。

*/

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/deatil/go-cryptobin/cryptobin/crypto"
)

// CaesarCipher implements a simple Caesar cipher encryption and decryption
func CaesarCipher(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}

// GetLetterByPosition returns the letter at the specified position in the alphabet
func GetLetterByPosition(position int) string {
	if position < 0 || position > 25 {
		return ""
	}
	return string(rune('A' + position))
}

// Json Crypto , make every Json was loaded into a encrypt one.
// remember only used in need area (e.g. crypto process.)

// CaesarCipher Code Design.

// How to work:
/*
	getTime Unix First, patched the third one - revert char 4 one , and do revert.
	always have 5 chars, if chars above 5 (means 6), it possibly came to 2286-11-21 01:46:40 ?
	Generate a number , as mixed one char.
	if this number is Disvisible, set x2 to this Caesar Code , else do nothing.
	Them, get the first number as char one, the last number as char 2
	get the position and mixed it as caesar code postion
	Last, get mid of three number, just caesar it with random code, the start of code is CaesarCode With A
*/

func JsonHashedData(k []byte) (passkey int, HandleTime int64, EncryptoData string) {
	randomCode, NewHashedCode, HandleTime := GenerateKey()
	return randomCode, HandleTime, CryptoHashedData(k, NewHashedCode)
}

func GenerateExistedKey(passkey int, unix int64) (HashedKey string) {
	makeHasherIntoStr := strconv.FormatInt(unix, 10)
	getCurrentStr := makeHasherIntoStr[2:(len(makeHasherIntoStr) - 3)]
	// Reverse the list
	reversedStr := ""
	for i := len(getCurrentStr) - 1; i >= 0; i-- {
		reversedStr += string(getCurrentStr[i])
	}
	getCurrentStr = reversedStr
	var MakeChartList []int
	for _, v := range getCurrentStr {
		MakeChartList = append(MakeChartList, int(v-'0'))
	}
	// generate a random number.
	randomCode := passkey
	RandIntoPosition := CaesarCipher("a", randomCode)
	// adjust if have reminder.
	// set vars
	var CharA string
	var CharB string
	if randomCode%2 == 0 {
		CharA = GetLetterByPosition(MakeChartList[0] * 2)
		CharB = GetLetterByPosition(MakeChartList[len(MakeChartList)-1] * 2)
	} else {
		CharA = GetLetterByPosition(MakeChartList[0])
		CharB = GetLetterByPosition(MakeChartList[len(MakeChartList)-1])
	}
	// set Char Code
	var NewHashedCode string
	// set last temp
	var NewTempLastCode string
	makeLeftList := MakeChartList[1:4]
	// make left leave.
	for _, i := range makeLeftList {
		NewTempLastCode += CaesarCipher(RandIntoPosition, i)
	}
	NewHashedCode = fmt.Sprintf("kvpk-%d%d%d%s%d%s%d%d%s", randomCode, MakeChartList[0], MakeChartList[1], CharA, MakeChartList[2], CharB, MakeChartList[3], MakeChartList[4], NewTempLastCode)
	return NewHashedCode
}

func GenerateKey() (passkey int, HashedKey string, HandleTime int64) {
	makeHasherIntoStr := strconv.FormatInt(time.Now().Unix(), 10)
	getCurrentStr := makeHasherIntoStr[2:(len(makeHasherIntoStr) - 3)]
	// Reverse the list
	reversedStr := ""
	for i := len(getCurrentStr) - 1; i >= 0; i-- {
		reversedStr += string(getCurrentStr[i])
	}
	getCurrentStr = reversedStr
	var MakeChartList []int
	for _, v := range getCurrentStr {
		MakeChartList = append(MakeChartList, int(v-'0'))
	}
	// generate a random number.
	randomCode := rand.Intn(8) + 1
	RandIntoPosition := CaesarCipher("a", randomCode)
	// adjust if have reminder.
	// set vars
	var CharA string
	var CharB string
	if randomCode%2 == 0 {
		CharA = GetLetterByPosition(MakeChartList[0] * 2)
		CharB = GetLetterByPosition(MakeChartList[len(MakeChartList)-1] * 2)
	} else {
		CharA = GetLetterByPosition(MakeChartList[0])
		CharB = GetLetterByPosition(MakeChartList[len(MakeChartList)-1])
	}
	// set Char Code
	var NewHashedCode string
	// set last temp
	var NewTempLastCode string
	makeLeftList := MakeChartList[1:4]
	// make left leave.
	for _, i := range makeLeftList {
		NewTempLastCode += CaesarCipher(RandIntoPosition, i)
	}
	NewHashedCode = fmt.Sprintf("kvpk-%d%d%d%s%d%s%d%d%s", randomCode, MakeChartList[0], MakeChartList[1], CharA, MakeChartList[2], CharB, MakeChartList[3], MakeChartList[4], NewTempLastCode)
	return randomCode, NewHashedCode, HandleTime
}

func DepressHashedData(passkey int, HandleTime int64, raw string) ([]byte, error) {
	HasedKey := GenerateExistedKey(passkey, HandleTime)
	host, err := base64.StdEncoding.DecodeString(CryptoDecryptData(raw, HasedKey))
	if err != nil {
		return nil, err
	}
	return host, err

}

func CryptoHashedData(content []byte, RawKey string) string {
	contentRaw := base64.StdEncoding.EncodeToString(content)
	return crypto.FromBase64String(contentRaw).SetKey(RawKey).Encrypt().ToBase64String()
}

func CryptoDecryptData(contentRawB64 string, RawKey string) string {
	return crypto.FromBase64String(contentRawB64).SetKey(RawKey).Decrypt().ToBase64String()
}
