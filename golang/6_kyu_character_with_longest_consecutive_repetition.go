/*

For a given string s find the character c (or C) with longest consecutive repetition and return:

type Result struct {
    C rune // character
    L int  // count
}
where l (or L) is the length of the repetition. If there are two or more characters with the same l return the first in order of appearance.

*/

package kata

func LongestRepetition(text string) Result {
	result := Result{}
	current := Result{}

	for _, r := range text {
		if current.C != r {
			updateResult(&current, &result)

			current.C = r
			current.L = 0
		}
		current.L++
	}

	updateResult(&current, &result)

	return result
}

func updateResult(current *Result, result*Result) {
	if current.L > result.L {
		result.C = current.C
		result.L = current.L
	}

	return
}
