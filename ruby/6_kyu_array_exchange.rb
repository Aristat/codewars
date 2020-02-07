=begin

Array Exchange and Reversing

It's time for some array exchange! The objective is simple: exchange the elements of two arrays in-place in a way that their new content is also reversed.

# before
my_array = ['a', 'b', 'c']
other_array = [1, 2, 3]

my_array.exchange_with!(other_array)

# after
my_array == [3, 2, 1]
other_array == ['c', 'b', 'a']

=end

class Array
  def exchange_with!(other_array)
    old_array = self.clone
    self.replace(other_array.reverse)
    other_array.replace(old_array.reverse)
  end
end
