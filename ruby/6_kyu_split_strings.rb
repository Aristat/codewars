# Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
# 
# Examples:
# 
# Solution("abc") //should return ["ab", "c_"]
# Solution("abcdef") //should return ["ab", "cd", "ef"]

def solution(str)
  str.split('').each_slice(2).map { |elements| elements.length == 2 ? elements.join('') : elements.first + "_"  }
end
