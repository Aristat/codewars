/*

Digital root is the recursive sum of all the digits in a number.

Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. This is only applicable to the natural numbers.

Examples
  16  -->  1 + 6 = 7
 	942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
	132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
	493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2

*/

package kata

func DigitalRoot(n int) int {
	var result int

	for {
		n = sumDigits(n)

		if n < 10 {
			result = n
			break
		}
	}

	return result
}

func sumDigits(number int) int {
	remainder := 0
	sumResult := 0

	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}

	return sumResult
}
