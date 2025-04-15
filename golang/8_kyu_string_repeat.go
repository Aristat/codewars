/*

Write a function called repeat_str which repeats the given string src exactly count times.

repeatStr(6, "I") // "IIIIII"
repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"

*/

package kata

func RepeatStr(repetitions int, value string) string {
	n := 0
	result := ""

	for n < repetitions {
		result += value
		n += 1
	}

	return result
}
