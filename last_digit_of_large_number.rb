DIGIT_TABLE = [
  [0],
  [1],
  [4,8,6,2],
  [9,7,1,3],
  [6,4],
  [5],
  [6],
  [9,3,1,7],
  [4,2,6,8],
  [1,9],
].freeze

def last_digit(number, exp)
  return 1 if exp == 0
  last_number = number % 10
  sequence = DIGIT_TABLE[last_number]
  
  remainder = exp % sequence.length
  remainder_for_last_digit = remainder == 0 ? sequence.length : remainder
  (last_number ** remainder_for_last_digit).digits.first
end
