=begin

k-Primes

A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity. For example:

k = 2  -->  4, 6, 9, 10, 14, 15, 21, 22, ...
k = 3  -->  8, 12, 18, 20, 27, 28, 30, ...
k = 5  -->  32, 48, 72, 80, 108, 112, ...
A natural number is thus prime if and only if it is 1-prime.

Task:

Complete the function count_Kprimes (or countKprimes, count-K-primes, kPrimes) which is given parameters k, start, end (or nd) and returns an array (or a list or a string depending on the language - see "Solution" and "Sample Tests") of the k-primes between start (inclusive) and end (inclusive).

=end

def count_Kprimes(k, start, nd)
  result = []

  (start..nd).each do |n|
    result << n if k_primes(n, k) == k
  end

  result
end

def puzzle(s)
  result = 0
  
  a = count_Kprimes(1, 1, s)
  b = count_Kprimes(3, 1, s)
  c = count_Kprimes(7, 1, s)

  return 0 if a == [] || b == [] || c == []

  a.each do |number_a|
    b.each do |number_b|
      c.each do |number_c|
        result += 1 if (number_a + number_b + number_c) == s
      end
    end
  end

  result
end

def k_primes(n, k)
  _kprimes = 0
  i = 2

  while i * i <= n
    while n % i == 0
      n /= i
      _kprimes += 1
    end

    break if _kprimes > k

    i += 1
  end

  _kprimes += 1 if n != 1

  _kprimes
end
