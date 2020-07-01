/*

Description:
Encrypt this!

You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter needs to be converted to its ASCII code.
The second letter needs to be switched with the last letter
Keepin' it simple: There are no special characters in input.

*/

package kata

import (
  "strconv"
  "strings"
)

func EncryptThis(text string) string {
  words := strings.Fields(text)
  values := []string{}
  for _, word := range words {
    sWord := strings.Split(word, "")
    wordLen := len(sWord)

    wordRunes := []rune(sWord[0])
    sWord[0] = strconv.Itoa(int(wordRunes[0]))
    if wordLen > 1 {
      sWord[1], sWord[wordLen-1] = sWord[wordLen-1], sWord[1]
    }

    values = append(values, strings.Join(sWord, ""))
  }

  return strings.Join(values, " ")
}
