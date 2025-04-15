/*

Given a lowercase string that has alphabetic characters only and no spaces, return the highest value of consonant substrings. Consonants are any letters of the alphabet except "aeiou".

We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.

For example, for the word "zodiacs", let's cross out the vowels. We get: "z o d ia cs"

-- The consonant substrings are: "z", "d" and "cs" and the values are z = 26, d = 4 and cs = 3 + 19 = 22. The highest is 26.
solve("zodiacs") = 26

For the word "strength", solve("strength") = 57
-- The consonant substrings are: "str" and "ngth" with values "str" = 19 + 20 + 18 = 57 and "ngth" = 14 + 7 + 20 + 8 = 49. The highest is 57.

*/

package kata

var consonantsRunes = []rune{'a', 'e', 'i', 'o', 'u'}

func solve(str string) int {
	strRunes := []rune(str)
	result := []int{0}

	for _, word := range strRunes {
		if contains(consonantsRunes, word) {
			if result[len(result)-1] != 0 {
				result = append(result, 0)
			}
			continue
		}

		result[len(result)-1] += toPosition(word)
	}

	return max(result)
}

func toPosition(i rune) int {
	return int(rune(i - 'a' + 1))
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
