/*

A format for expressing an ordered list of integers is to use a comma separated list of either

individual integers
or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'.
The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers.
For example ("12, 13, 15-17")
Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.

*/

package kata

import "strconv"

func Solution(list []int) string {
	listLen := len(list)
	var counter int
	var numbers []string
	var result string

	for i := 0; i < listLen; i++ {
		counter = 0
		start := strconv.Itoa(list[i])

		for i != listLen-1 && (list[i+1]-list[i]) == 1 {
			counter++
			i++
		}

		if counter > 0 {
			if counter > 1 {
				i, _ := strconv.Atoi(start)
				start += "-" + strconv.Itoa(i+counter)
			} else {
				i--
			}
		}

		numbers = append(numbers, start)
	}

	numbersLen := len(numbers)
	for l, n := range numbers {
		result += n
		if l != numbersLen-1 {
			result += ","
		}
	}

	return result
}
