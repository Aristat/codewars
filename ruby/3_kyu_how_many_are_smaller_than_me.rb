=begin

This is a hard version of How many are smaller than me?. If you have troubles solving this one, have a look at the easier kata first.

Write

function smaller(arr)
that given an array arr, you have to return the amount of numbers that are smaller than arr[i] to the right.

For example:

smaller([5, 4, 3, 2, 1]) === [4, 3, 2, 1, 0]
smaller(1, 2, 0) === [1, 1, 0]
  
=end

class Node
  attr_accessor :duplicates, :value, :sum, :right, :left

  def initialize(v, s)
    @duplicates = 1
    @value = v
    @sum = s
  end
end

def insert(num, index, node, result, previous_sum)
  if node === nil
    node = Node.new(num, 0)
    result[index] = previous_sum
  elsif node.value == num
    node.duplicates += 1
    result[index] = previous_sum + node.sum
  elsif node.value > num
    node.sum += 1
    node.left = insert(num, index, node.left, result, previous_sum)
  else
    node.right = insert(num, index, node.right, result, previous_sum + node.duplicates + node.sum)
  end

  node
end

def smaller(arr)
  result = []
  root = nil

  arr.to_enum.with_index.reverse_each do |v, index|
    root = insert(v, index, root, result, 0)
  end

  result
end
