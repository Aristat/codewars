def done_or_not(board)
  if check_sum(board) && check_regions(board)
    'Finished!'
  else
    'Try again!'
  end
end

def check_sum(board)
  board.transpose.all? {|a| a.uniq.length == 9 && a.inject(0, :+) == 45} && 
    board.all? { |a| a.uniq.length == 9 && a.inject(0, :+) == 45}
end

def check_regions(board)
  slices = board.each_slice(3).to_a
  regions = slices.map {|s| s.transpose.reduce(:+).each_slice(9).to_a}.reduce(:+)
  regions.map{|e| e.reduce(:+) == 45}.reduce(:&)
end
