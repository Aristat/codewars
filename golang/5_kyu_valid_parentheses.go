/*

Write a function called that takes a string of parentheses, and determines if the order of the parentheses is valid. The function should return true if the string is valid, and false if it's invalid.

Examples
"()"              =>  true
")(()))"          =>  false
"("               =>  false
"(())((()())())"  =>  true

*/

package kata

func ValidParentheses(parens string) bool {
  stick := []rune{}
  
  for _, character := range(parens) {
    if character == '(' {
      stick = append(stick, character)
    } else if character == ')' {
      if len(stick) < 1 {
        return false
      }
      stick = stick[:len(stick)-1]
    }
  }

  if len(stick) == 0 {
    return true
  }
  
  return false
}
