/*

The input is a string str of digits. Cut the string into chunks (a chunk here is a substring of the initial string) of size sz (ignore the last chunk if its size is less than sz).

If a chunk represents an integer such as the sum of the cubes of its digits is divisible by 2, reverse that chunk; otherwise rotate it to the left by one position.
Put together these modified chunks and return the result as a string.

If

sz is <= 0 or if str is empty return ""
sz is greater (>) than the length of str it is impossible to take a chunk of size sz hence return "".
Examples:
revrot("123456987654", 6) --> "234561876549"
revrot("123456987653", 6) --> "234561356789"
revrot("66443875", 4) --> "44668753"
revrot("66443875", 8) --> "64438756"
revrot("664438769", 8) --> "67834466"
revrot("123456779", 8) --> "23456771"
revrot("", 8) --> ""
revrot("123456779", 0) --> ""
revrot("563000655734469485", 4) --> "0365065073456944"

*/

package kata

import (
	"math"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if n <= 0 || s == "" || n > len(s) {
		return ""
	}

	words := strings.Split(s, "")
	chunksStrings := []string{}

	for len(words) >= n {
		chunk := append([]string(nil), words[:n]...)
		words = words[n:]

		sum := 0
		for _, s := range chunk {
			n, _ := strconv.Atoi(s)
			sum += int(math.Pow(float64(n), float64(3)))
		}

		if sum%2 != 0 {
			number := chunk[0]
			chunk = chunk[1:]
			chunk = append(chunk, number)
		} else {
			for i := len(chunk)/2 - 1; i >= 0; i-- {
				opp := len(chunk) - 1 - i
				chunk[i], chunk[opp] = chunk[opp], chunk[i]
			}
		}

		chunksStrings = append(chunksStrings, strings.Join(chunk[:], ""))
	}

	return strings.Join(chunksStrings[:], "")
}
