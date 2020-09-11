/*

The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

Examples
"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))(("

*/

package kata

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	dictionary := make(map[rune]int)
	result := ""

	for _, char := range word {
		dictionary[char] += 1
	}

	for _, char := range word {
		if dictionary[char] > 1 {
			result = result + ")"
		} else {
			result = result + "("
		}
	}

	return result
}
