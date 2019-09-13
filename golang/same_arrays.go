package kata

import (
  "sort"
  "reflect"
)

func Comp(array1 []int, array2 []int) bool {
  if array1 == nil || array2 == nil || len(array1) != len(array2) {
    return false
  }
  
  newArray := []int{}

  for _, number := range array1 {
    newArray = append(newArray, number*number) 
  }  
  
  sort.Ints(newArray)
  sort.Ints(array2)
  return reflect.DeepEqual(newArray, array2)
}
