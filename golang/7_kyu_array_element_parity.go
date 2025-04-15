/*

In this Kata, you will be given an array of integers whose elements have both a negative and a positive value, except for one integer that is either only negative or only positive. Your task will be to find that integer.

Examples:

[1, -1, 2, -2, 3] => 3

3 has no matching negative appearance

[-3, 1, 2, 3, -1, -4, -2] => -4

-4 has no matching positive appearance

[1, -1, 2, -2, 3, 3] => 3

(the only-positive or only-negative integer may appear more than once)

*/

package kata

import "math"

func Solve(arr []int) int {
	valuesMap := map[int][]int{}

	for _, k := range arr {
		key := int(math.Abs(float64(k)))
		valuesMap[key] = appendIfMissing(valuesMap[key], k)
	}

	number := 0
	for _, v := range valuesMap {
		if len(v) == 1 {
			number = v[0]
			break
		}
	}

	return number
}

func appendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
