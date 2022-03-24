package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
)

func CheckContainsTime(s string) bool {
	re := regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, etc
	return re.MatchString(s)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[index.Int64()]
	}
	return string(b)
}

func GenerateFakesCode() string {
	code := fmt.Sprintf("FAKES-%v", stringWithCharset(10, "1234567890ABCDEFG"))
	return code
}

func RandomNumber(length int) string {
	return stringWithCharset(length, "1234567890")
}
