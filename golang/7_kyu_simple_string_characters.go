/*

In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3].
--the order is: uppercase letters, lowercase, numbers and special characters.

 */

package kata

func Solve(s string) []int {
	uppercase := 0
	lowercase := 0
	numbers := 0
	special := 0

	for _, c := range s {
		if c >= 65 && c <= 90 {
			uppercase += 1
		} else if c >= 97 && c <= 122 {
			lowercase += 1
		} else if c >= 48 && c <= 57 {
			numbers += 1
		} else {
			special += 1
		}
	}

	return []int{uppercase, lowercase, numbers, special}
}
