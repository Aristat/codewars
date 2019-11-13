/*

Given an array, find the int that appears an odd number of times.
There will always be only one integer that appears an odd number of times.

*/

package kata

func FindOdd(seq []int) int {
  var result int
  duplicates := make(map[int]int)

  for _, item := range seq {
    _, exist := duplicates[item]

    if exist {
      duplicates[item] += 1
    } else {
      duplicates[item] = 1
    }
  }

  for number, item := range duplicates {
    if item%2 != 0 {
      result = number
      break
    }
  }
  
  return result
}
