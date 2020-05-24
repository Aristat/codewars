=begin

Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the row sums of this triangle from the row index (starting at index 1) e.g.:

row_sum_odd_numbers(1); # 1
row_sum_odd_numbers(2); # 3 + 5 = 8
  
=end

def row_sum_odd_numbers(n)
  start = n ** 2 - (n - 1)
  finish = start + n * 2 - 2

  sum = start
  while start < finish
    start += 2
    sum += start
  end

  return sum
end
