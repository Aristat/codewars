=begin
  
Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

The input will be a lowercase string with no spaces.

=end

def capitalize(s)
  even_result = ""
  odd_result = ""
 
  s.chars.each_with_index do |i, index|
    if index.even?
      even_result << i.upcase
      odd_result << i
    else
      even_result << i
      odd_result << i.upcase
    end
  end
  
  [even_result, odd_result]
end
