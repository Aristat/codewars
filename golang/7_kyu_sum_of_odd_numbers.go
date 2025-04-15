/*

Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the row sums of this triangle from the row index (starting at index 1) e.g.:

rowSumOddNumbers(1); // 1
rowSumOddNumbers(2); // 3 + 5 = 8

*/

package kata

import (
	"math"
)

func RowSumOddNumbers(n int) int {
	start := int(math.Pow(float64(n), 2)) - (n - 1)
	finish := start + n*2 - 2
	sum := start

	for start < finish {
		start += 2
		sum += start
	}

	return sum
}
