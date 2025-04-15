/*

Description:

Input : an array of integers.
Output : this array, but sorted in such a way that there are two wings:

the left wing with numbers decreasing,
the right wing with numbers increasing.

the two wings have the same length. If the length of the array is odd the wings are around the bottom, if the length is even the bottom is considered to be part of the right wing.
each integer l of the left wing must be greater or equal to its counterpart r in the right wing, the difference l - r being as small as possible. In other words the right wing must be nearly as steep as the left wing.

*/

package kata

import "sort"

func MakeValley(arr []int) []int {
	arrLen := len(arr)

	var output []int
	if arrLen == 0 {
		output = []int{}
		return output
	}

	sort.Ints(arr)

	var startIndex int
	if arrLen%2 == 0 {
		output = []int{arr[0]}
		startIndex = 1
	} else {
		output = []int{}
		startIndex = 0
	}

	for index, element := range arr[startIndex:] {
		if index%2 == 0 {
			output = append([]int{element}, output...)
		} else {
			output = append(output, element)
		}
	}

	return output
}
