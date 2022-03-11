=begin

Given an array of numbers and an index, return the index of the least number larger than the element at the given index,
or -1 if there is no such index ( or, where applicable, Nothing or a similarly empty value ).

Notes
Multiple correct answers may be possible. In this case, return any one of them.
The given index will be inside the given array.
The given array will, therefore, never be empty.

=end

def least_larger(a, i)
  n = a[i]
  r = nil

  a.each_with_index do |number, index|
    next if number <= n

    if r.nil?
      r = index
      next
    end

    r = a[r] > number ? index : r
  end

  r.nil? ? -1 : r
end
