package isPalindrom

import (
	"strings"
)

// Needed to use Split

func IsPalindrom(str string) bool {
	arrayStr := []string{}
	strSplit := strings.Split(str, "")

	for i := len(strSplit) - 1; i >= 0; i-- {
		arrayStr = append(arrayStr, strSplit[i])
	}
	revesedStr := strings.Join(arrayStr, "")

	return revesedStr == str

}
