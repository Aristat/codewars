/*

In this Kata, you will be given a string that may have mixed uppercase and lowercase letters and your task is to convert that string to either lowercase only or uppercase only based on:

make as few changes as possible.
if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
For example:

solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
solve("coDE") = "code". Upper == lowercase. Change all to lowercase.

*/

package kata

import (
	"strings"
	"unicode"
)

func solve(str string) string {
	countLower := 0
	countUpper := 0

	for _, char := range str {
		if unicode.IsLower(char) {
			countLower += 1
		} else {
			countUpper += 1
		}
	}

	if countLower >= countUpper {
		str = strings.ToLower(str)
	} else {
		str = strings.ToUpper(str)
	}

	return str
}
