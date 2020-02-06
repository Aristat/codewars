=begin

Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

Return your answer as a number.

=end

def sum_mix(x)
  x.map(&:to_i).sum
end
