# A *Hamming number* is a positive integer of the form 2i3j5k, for some non-negative integers i, j, and k.

# Write a function that computes the nth smallest Hamming number.

# Specifically:

# The first smallest Hamming number is 1 = 203050
# The second smallest Hamming number is 2 = 213050
# The third smallest Hamming number is 3 = 203150
# The fourth smallest Hamming number is 4 = 223050
# The fifth smallest Hamming number is 5 = 203051
# The 20 smallest Hamming numbers are given in example test fixture.

# Your code should be able to compute all of the smallest 5,000 (Clojure: 2000) Hamming numbers without timing out.

def hamming(n)
  count = 1
  hamming_numbers = [1]
  
  i = 0
  j = 0
  k = 0
  
  while count < n
    hamming_numbers[count] = [hamming_numbers[i] * 2, hamming_numbers[j] * 3, hamming_numbers[k] * 5].min

    i += 1 if (hamming_numbers[count] == hamming_numbers[i] * 2) 
    j += 1 if (hamming_numbers[count] == hamming_numbers[j] * 3)
    k += 1 if (hamming_numbers[count] == hamming_numbers[k] * 5)

    count += 1
  end
  
  hamming_numbers.last
end