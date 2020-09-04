/*

Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number.
You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

 */

package kata

import (
	"strconv"
)

func CountBits(n uint) int {
	var one rune = 49
	count := 0

	bs := strconv.FormatInt(int64(n), 2)

	for _, char := range bs {
		if char == one {
			count += 1
		}
	}

	return count
}
