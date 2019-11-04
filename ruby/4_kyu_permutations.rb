=begin

In this kata you have to create all permutations of an input string and remove duplicates, if present. 
This means, you have to shuffle all letters from the input in all possible orders.

Examples:

permutations('a'); # ['a']
permutations('ab'); # ['ab', 'ba']
permutations('aabb'); # ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa']
The order of the permutations doesn't matter.
  
=end

def permutations(string)
  _permutations(string.split('')).uniq.sort
end

def _permutations(array, i=0, result = [])
  if i == array.size
    result << array.join('')
  end

  (i..array.size-1).each do |j|
    array[i], array[j] = array[j], array[i]
    _permutations(array, i+1, result)
    array[i], array[j] = array[j], array[i]
  end

  result
end
