/*

The drawing shows 6 squares the sides of which have a length of 1, 1, 2, 3, 5, 8.
It's easy to see that the sum of the perimeters of these squares is : 4 * (1 + 1 + 2 + 3 + 5 + 8) = 4 * 20 = 80

The function perimeter has for parameter n where n + 1 is the number of squares (they are numbered from 0 to n) and returns the total perimeter of all the squares.

perimeter(5)  should return 80
perimeter(7)  should return 216

*/

package kata

func Perimeter(n int) int {
	fn := make(map[int]int)

	for i := 1; i <= n+1; i++ {
		if i <= 2 {
			fn[i] = 1
		} else {
			fn[i] = fn[i-1] + fn[i-2]
		}
	}

	sum := 0
	for _, value := range fn {
		sum += value
	}

	return sum * 4
}
