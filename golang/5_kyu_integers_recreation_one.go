/*

Divisors of 42 are : 1, 2, 3, 6, 7, 14, 21, 42. These divisors squared are: 1, 4, 9, 36, 49, 196, 441, 1764. The sum of the squared divisors is 2500 which is 50 * 50, a square!

Given two integers m, n (1 <= m <= n) we want to find all integers between m and n whose sum of squared divisors is itself a square. 42 is such a number.

The result will be an array of arrays or of tuples (in C an array of Pair) or a string, each subarray having two elements, first the number whose squared divisors is a square and then the sum of the squared divisors.

#Examples:

list_squared(1, 250) --> [[1, 1], [42, 2500], [246, 84100]]
list_squared(42, 250) --> [[42, 2500], [246, 84100]]
The form of the examples may change according to the language, see Example Tests: for more details.

Note

In Fortran - as in any other language - the returned string is not permitted to contain any redundant trailing whitespace: you can use dynamically allocated character strings.

*/

package kata

import (
	"math"
)

func ListSquared(m, n int) [][]int {
	matches := [][]int{}

	for i := m; i <= n; i++ {
		divisors := []int{}

		for j := 1; j <= i/2; j++ {
			if i%j != 0 {
				continue
			}

			divisors = append(divisors, j)
		}

		divisors = append(divisors, i)

		sum := 0
		for _, value := range divisors {
			sum += value * value
		}

		if math.Trunc(math.Sqrt(float64(sum))) == math.Sqrt(float64(sum)) {
			matches = append(matches, []int{i, sum})
		}
	}

	return matches
}
