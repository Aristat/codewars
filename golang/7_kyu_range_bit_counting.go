/*

Task
You are given two numbers a and b where 0 ≤ a ≤ b. Imagine you construct an array of all the integers from a to b inclusive. You need to count the number of 1s in the binary representations of all the numbers in the array.

Example
For a = 2 and b = 7, the output should be 11

Given a = 2 and b = 7 the array is: [2, 3, 4, 5, 6, 7]. Converting the numbers to binary, we get [10, 11, 100, 101, 110, 111], which contains 1 + 2 + 1 + 2 + 2 + 3 = 11 1s.

Input/Output
[input] integer a

Constraints: 0 ≤ a ≤ b.

[input] integer b

Constraints: a ≤ b ≤ 100.

[output] an integer

*/

package kata

import "fmt"

func RangeBitCount(a, b int) int {
	var asciiNum int32 = 49
	count := 0
	arr := makeRange(a, b)

	for _, i := range arr {
		binary := fmt.Sprintf("%b", i)

		for _, num := range binary {
			if num == asciiNum {
				count += 1
			}
		}
	}

	return count
}

func makeRange(min, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}
