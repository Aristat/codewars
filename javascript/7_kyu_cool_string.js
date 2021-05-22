/*

Let's call a string cool if it is formed only by Latin letters and no two lowercase and no two uppercase letters are in adjacent positions. 
Given a string, check if it is cool.

*/

function coolString(s) {
  let chars = [...s].map(char => char.charCodeAt(0))
  let lChars = chars.length
  let cool = true

  for (let i = 0; i < lChars; i++) {
    if (i === lChars - 1 && (isUppercase(chars[i]) || isLowercase(chars[i])))
      break;

    if (isUppercase(chars[i]) && !isUppercase(chars[i + 1]))
      continue;

    if (isLowercase(chars[i]) && !isLowercase(chars[i + 1]))
      continue;

    cool = false
    break;
  }

  return cool
}

function isUppercase(c) {
  return c >= 65 && c <= 90
}

function isLowercase(c) {
  return c >= 97 && c <= 122
}
