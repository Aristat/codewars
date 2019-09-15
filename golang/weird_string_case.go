package kata

import "strings"

func toWeirdCase(str string) string {
  newString := ""
  words := strings.Fields(str)

  for index, word := range words {
    for index, character := range word {
      if index % 2 == 0 {
        newString += strings.ToUpper(string(character))
      } else {
        newString += strings.ToLower(string(character))
      }
    }
    
    if index < len(words) - 1 {
      newString += " "
    }
  }
  
  return newString
}
