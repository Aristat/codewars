/*

Rules are: (1) the letters are adjacent in the English alphabet, and (2) each letter occurs only once.

For example:
solve("abc") = True, because it contains a,b,c
solve("abd") = False, because a, b, d are not consecutive/adjacent in the alphabet, and c is missing.
solve("dabc") = True, because it contains a, b, c, d
solve("abbc") = False, because b does not occur once.
solve("v") = True

 */

package kata

func Solve(s string) bool {
	var max int
	var min int
	var count int

	for _, char := range s {
		intChart := int(char)

		if intChart > max || max == 0 {
			max = intChart
		}

		if intChart < min || min == 0 {
			min = intChart
		}

		count += 1
	}

	if (max - min) == (count - 1) {
		return true
	} else {
		return false
	}
}
