/*

In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

For example:

solve("our code") = "edo cruo"
-- Normal reversal without spaces is "edocruo".
-- However, there is a space at index 3, so the string becomes "edo cruo"

solve("your code rocks") = "skco redo cruoy".
solve("codewars") = "srawedoc"

*/

package kata

func Solve(s string) string {
	var reverseChars string

	for _, e := range s {
		if e == 32 {
			continue
		}

		reverseChars = string(e) + reverseChars
	}

	result := ""
	charIndex := 0

	for _, e := range s {
		if e == 32 {
			result = result + " "
			continue
		}

		result = result + string(reverseChars[charIndex])
		charIndex += 1
	}

	return result
}

