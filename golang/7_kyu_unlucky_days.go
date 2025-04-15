/*

Friday 13th or Black Friday is considered as unlucky day. Calculate how many unlucky days are in the given year.

Find the number of Friday 13th in the given year.

Input: Year as an integer.

Output: Number of Black Fridays in the year as an integer.

Examples:

unluckyDays(2015) == 3
unluckyDays(1986) == 1

*/

package kata

import (
	"time"
)

func UnluckyDays(year int) int {
	count := 0

	for i := 1; i <= 12; i++ {
		weekday := time.Date(year, time.Month(i), 13, 0, 0, 0, 0, time.UTC)

		if weekday.Weekday().String() == "Friday" {
			count += 1
		}
	}

	return count
}
