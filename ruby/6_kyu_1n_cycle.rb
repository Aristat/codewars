=begin

Let be n an integer prime with 10 e.g. 7.

1/7 = 0.142857 142857 142857 ....

We see that the decimal part has a cycle: 142857. The length of this cycle is 6. In the same way:

1/11 = 0.09 09 09 .... Cycle length is 2.

Task
Given an integer n (n > 1), the function cycle(n) returns the length of the cycle if n and 10 are coprimes, otherwise returns -1.

=end

def cycle(n)
  return -1 if n % 2 == 0 || n % 5 == 0

  res = 10 % n
  c = 1
  
  while (res != 1) do
    res = res * 10 % n
    c += 1
  end
  
  c
end
