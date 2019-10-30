=begin

Counting Duplicates

Count the number of Duplicates

Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits 
that occur more than once in the input string. 

The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

=end

def duplicate_count(text)
  text.downcase.chars.group_by{ |e| e }.select { |k, v| v.size > 1 }.map(&:first).length
end
