=begin

Unique In Order

Implement the function unique_in_order which takes as argument a sequence and returns a 
list of items without any elements with the same value next to each other and preserving the original order of elements.

=end

def unique_in_order(iterable)
  iterable= iterable.split('') if iterable.is_a?(String)
  new_array = []
  last_element = nil
  iterable.each  do |i|
    next if i == last_element
    last_element = i
    new_array << i
  end
  
  new_array
end
