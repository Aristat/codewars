/*

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, and u as vowels for this Kata.

The input string will only consist of lower case letters and/or spaces.

*/

package kata

func GetCount(str string) (count int) {
	m := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}

	for _, item := range str {
		_, ok := m[string(item)]
		if ok {
			count += 1
		}

	}

	return count
}
