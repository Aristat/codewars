/*

In this kata, you will sort elements in an array by decreasing frequency of elements. If two elements have the same frequency, sort them by increasing value.

solve([2,3,5,3,7,9,5,3,7]) = [3,3,3,5,5,7,7,2,9]
-- We sort by highest frequency to lowest frequency.
-- If two elements have same frequency, we sort by increasing value.

*/

package kata

import (
	"sort"
)

func Solve(arr []int) []int {
	hashMap := make(map[int]int)

	for _, v := range arr {
		hashMap[v]++
	}

	keys := make([]int, 0, len(hashMap))
	for k := range hashMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if hashMap[keys[i]] == hashMap[keys[j]] {
			return keys[i] < keys[j]
		}
		return hashMap[keys[i]] > hashMap[keys[j]]
	})

	result := make([]int, 0, len(keys))
	for _, v := range keys {
		for i := 0; i < hashMap[v]; i++ {
			result = append(result, v)
		}
	}

	return result
}
