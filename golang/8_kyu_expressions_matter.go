/*

Task
Given three integers a ,b ,c, return the largest number obtained after inserting the following operators and brackets: +, *, ().

Consider an Example :
With the numbers are 1, 2 and 3 , here are some ways of placing signs and brackets:

1 * (2 + 3) = 5
1 * 2 * 3 = 6
1 + 2 * 3 = 7
(1 + 2) * 3 = 9

So the maximum value that you can obtain is 9.

*/

package kata

func ExpressionMatter(a int, b int, c int) int {
	expressions := []int{(a + b) * c, a * (b + c), a + b + c, a * b * c}
	max_value := 0

	for i, e := range expressions {
		if i == 0 || e > max_value {
			max_value = e
		}
	}

	return max_value
}
