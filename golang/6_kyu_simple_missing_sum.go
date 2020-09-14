/*

In this Kata, we will calculate the minumum positive number that is not a possible sum from a list of positive integers.

solve([1,2,8,7]) = 4 => we can get 1, 2, 3 (from 1+2), but we cannot get 4. 4 is the minimum number not possible from the list.
solve([4,1,2,3,12]) = 11. We can get 1, 2, 3, 4, 4+1=5, 4+2=6,4+3=7,4+3+1=8,4+3+2=9,4+3+2+1=10. But not 11.
solve([2,3,2,3,4,2,12,3]) = 1. We cannot get 1.

*/

package kata

import "sort"

func Solve(arr []int) int {
	sort.Ints(arr)

	result := 1

	for i := 0; i < len(arr) && arr[i] <= result; i++ {
		result = result + arr[i]
	}

	return result
}
