/*

Do you know how to make a spiral? Let's test it!

Classic definition: A spiral is a curve which emanates from a central point, getting progressively farther away as it revolves around the point.

Your objective is to complete a function createSpiral(N) that receives an integer N and returns an NxN two-dimensional array with numbers 1 through NxN represented as a clockwise spiral.

Return an empty array if N < 1 or N is not int / number

*/

package kata

func CreateSpiral(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	for num <= n*n {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
