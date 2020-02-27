/*

Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

Examples:

Solution("abc") //should return ["ab", "c_"]
Solution("abcdef") //should return ["ab", "cd", "ef"]

*/

package kata

func Solution(str string) []string {
  result := []string{}
  characters := ""

  for i, r := range str {
    characters = characters + string(r)
    if (i+1)%2 == 0 {
      result = append(result, characters)
      characters = ""
    }
  }

  if characters != "" {
    characters += "_"
    result = append(result, characters)
  }
  
  return result
}
