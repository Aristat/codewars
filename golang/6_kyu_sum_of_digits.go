package main

import (
	"fmt"
)

func sumDigits(number int) int {  
 remainder := 0  
 sumResult := 0  
 for number != 0 {  
  remainder = number % 10  
  sumResult += remainder  
  number = number / 10  
 }  
 return sumResult  
}

func main() {
	n := 195
	var result int

	for {
			n = sumDigits(n)

			if n < 10 {
				result = n
				break
			}
	}

	fmt.Println(result)
}
