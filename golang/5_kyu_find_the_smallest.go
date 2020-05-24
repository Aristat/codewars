/*

You have a positive number n consisting of digits. You can do at most one operation: Choosing the index of a digit in the number, remove this digit at that index and insert it back to another or at the same place in the number in order to find the smallest number you can get.

#Task: Return an array or a tuple or a string depending on the language (see "Sample Tests") with

1) the smallest number you got
2) the index i of the digit d you took, i as small as possible
3) the index j (as small as possible) where you insert this digit d to have the smallest number.
Example:

smallest(261235) --> [126235, 2, 0] or (126235, 2, 0) or "126235, 2, 0"
126235 is the smallest number gotten by taking 1 at index 2 and putting it at index 0

smallest(209917) --> [29917, 0, 1] or ...

[29917, 1, 0] could be a solution too but index `i` in [29917, 1, 0] is greater than 
index `i` in [29917, 0, 1].
29917 is the smallest number gotten by taking 2 at index 0 and putting it at index 1 which gave 029917 which is the number 29917.

smallest(1000000) --> [1, 0, 6] or ...

*/

package kata

import (
  "sort"
)

func Smallest(n int64) []int64 {
  numbers := int64ToSlice(n)
  allVariants := []int64{}
  hashVariants := map[int64][2]int{}

  for i, _ := range numbers {
    for j := 0; j < len(numbers); j++ {
      if i == j {
        intVariant := sliceToInt64(numbers)
        allVariants = append(allVariants, intVariant)

        setIntoHashVariants(hashVariants, intVariant, i, j)
      } else {
        variant := append([]int64{}, numbers...)
        element := numbers[i]                                                      // copy element
        variant = append(variant[:i], variant[i+1:]...)                            // generate variant without one element
        variant = append(variant[:j], append([]int64{element}, variant[j:]...)...) // insert element into slice

        intVariant := sliceToInt64(variant)
        allVariants = append(allVariants, intVariant)

        setIntoHashVariants(hashVariants, intVariant, i, j)
      }
    }
  }

  intMin := minIntSlice(allVariants)

  return []int64{intMin, int64(hashVariants[intMin][0]), int64(hashVariants[intMin][1])}
}

// insert into hash for fast finding from,to indexes
func setIntoHashVariants(hashVariants map[int64][2]int, variant int64, i int, j int) {
  oldValue, ok := hashVariants[variant]
  if ok {
    if oldValue[0] > i {
      hashVariants[variant] = [2]int{i, j}
    }

  } else {
    hashVariants[variant] = [2]int{i, j}
  }
}

// min element in slice
func minIntSlice(v []int64) int64 {
  sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
  return v[0]
}

// convert number to slice -> 1357 - [1,3,5,7]
func int64ToSlice(n int64) []int64 {
  s := []int64{}

  for n > 0 {
    mod := n % 10
    s = append([]int64{mod}, s...)

    n = n / 10
  }

  return s
}

// convert slice to number -> [1,3,5,7] - 1357
func sliceToInt64(s []int64) int64 {
  res := int64(0)
  op := int64(1)

  for i := int64(len(s)) - 1; i >= 0; i-- {
    res += s[i] * op
    op *= 10
  }

  return res
}
