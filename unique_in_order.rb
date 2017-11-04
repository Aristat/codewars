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
