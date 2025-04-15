/*

John has invited some friends. His list is:

s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";
Could you make a program that

makes this string uppercase
gives it sorted in alphabetical order by last name.
When the last names are the same, sort them by first name. Last name and first name of a guest come in the result between parentheses separated by a comma.

So the result of function meeting(s) will be:

"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
It can happen that in two distinct families with the same family name two people have the same first name too.

*/

package kata

import (
	"sort"
	"strings"
)

func Meeting(s string) string {
	firstAndLastNames := strings.Split(s, ";")

	var dict [][]string

	for _, v := range firstAndLastNames {
		var fLName []string
		values := strings.Split(v, ":")
		fLName = append(fLName, strings.ToUpper(values[1]))
		fLName = append(fLName, strings.ToUpper(values[0]))
		dict = append(dict, fLName)
	}

	sort.Slice(dict, func(i, j int) bool {
		if dict[i][0] < dict[j][0] {
			return true
		}
		if dict[i][0] > dict[j][0] {
			return false
		}
		return dict[i][1] < dict[j][1]
	})

	result := ""

	for _, v := range dict {
		result = result + "(" + v[0] + ", " + v[1] + ")"
	}

	return result
}
