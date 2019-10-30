=begin

In mathematics, a Diophantine equation is a polynomial equation, usually with two or more unknowns, such that only the integer solutions are sought or studied.

In this kata we want to find all integers x, y (x >= 0, y >= 0) solutions of a diophantine equation of the form:

x2 - 4 * y2 = n
(where the unknowns are x and y, and n is a given positive number) in decreasing order of the positive xi.

If there is no solution return [] or "[]" or "". (See "RUN SAMPLE TESTS" for examples of returns).

Hint:
x2 - 4 * y2 = (x - 2*y) * (x + 2*y)

=end

def sol_equa(n)
  result = []
  
  (1..(Math.sqrt(n)+1)).each do |i|
    if n % i != 0
      next
    end
  
    j = n / i
    y = (j - i) / 4
    x = i + 2 * y
  
    if x >= 0 and y >= 0 and (j == x + 2 * y) and (i == x - 2 * y)
      result << [x, y]
    end
  end
  
  result
end
