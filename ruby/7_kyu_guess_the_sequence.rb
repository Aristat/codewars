=begin

You have read the title: you must guess a sequence. It will have something to do with the number given.

Example
x = 16

result = [1, 10, 11, 12, 13, 14, 15, 16, 2, 3, 4, 5, 6, 7, 8, 9]
Good luck!

=end

def sequence(x)
  return [] if x == 0
  return (1..x).to_a if x <= 9

  n, d = x.divmod(10)
  r = []

  1.upto(n).each do |i|
    e = i * 10
    step = i == n ? d : 9

    r = r | [i] | (e..e + step).to_a
  end

  r | (1..x).to_a
end
