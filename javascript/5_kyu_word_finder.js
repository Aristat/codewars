/*
  In this kata you have to extend the dictionary with a method, that returns a list of words matching a pattern. 
  This pattern may contain letters (lowercase) and placeholders ("?"). A placeholder stands for exactly one arbitrary letter.

  Example:

  var fruits = new Dictionary(['banana', 'apple', 'papaya', 'cherry']);
  fruits.getMatchingWords('lemon');     // must return []
  fruits.getMatchingWords('cherr??');   // must return []
  fruits.getMatchingWords('?a?a?a');    // must return ['banana', 'papaya']
  fruits.getMatchingWords('??????');    // must return ['banana', 'papaya', 'cherry']
  Additional Notes:

  the words and patterns are all lowercase
  the order of the returned words doesn't matter
*/

function Dictionary(words) {
  this.words = words;
}

Dictionary.prototype.getMatchingWords = function(pattern) {
  result = []
  patternLength = pattern.length
  
  this.words.forEach(function (word) {
    if (word.length != patternLength) {
      return
    }
  
    correct = true
    
    for (var i = 0; i < patternLength; i++) {
      char = pattern.charAt(i)
      if (char == '?') continue
  
  
      if (word[i] != char) {
        correct = false
        break
      }
    }
  
    if (correct) {
      result.push(word)
    }

  });
  
  return result
}
