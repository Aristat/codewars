=begin

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Example:

high_and_low("1 2 3 4 5")  # return "5 1"
high_and_low("1 2 -3 4 5") # return "5 -3"
high_and_low("1 9 3 4 -5") # return "9 -5"
Notes:

All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.

=end

def high_and_low(numbers)
  max = nil
  min = nil
  numbers.split(' ').each do |i|
    max = i.to_i if max.nil? || i.to_i > max
    min = i.to_i if min.nil? || i.to_i < min
  end
  
  "#{max} #{min}"
end
