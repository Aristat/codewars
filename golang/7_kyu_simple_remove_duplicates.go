/*

In this Kata, you will remove the left-most duplicates from a list of integers and return the result.

// Remove the 3's at indices 0 and 3
// followed by removing a 4 at index 1
solve([3, 4, 4, 3, 6, 3]); // => [4, 6, 3]

*/

package kata

func Solve(arr []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		if encountered[arr[i]] != true {
			encountered[arr[i]] = true
			result = append([]int{arr[i]}, result...)
		}
	}

	return result
}
