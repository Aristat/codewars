=begin

Your task, is to create a NxN spiral with a given size.

For example, spiral with size 5 should look like this:

00000
....0
000.0
0...0
00000

Because of the edge-cases for tiny spirals, the size will be at least 5.
General rule-of-a-thumb is, that the snake made with '1' cannot touch to itself.

=end

def can_move(map, i, j, di, dj)
  n = map.length
  i += di
  j += dj
  
  if (i < 0 or i >= n or j < 0 or j >= n) || map[i][j] == 1
    return false
  end 
    
  i += di
  j += dj
  
  if i < 0 or i >= n or j < 0 or j >= n
    return true
  end 
  
  if map[i][j] == 1
    return false
  end  

  true
end

def spiralize(size)
  if size == 0
    return []
  elsif size == 1
    return [[1]]
  elsif size == 2
    return [[1,1],[0,1]]
  end
    
  spiral = Array.new(size) { Array.new(size, 0) }

  i, j = 0, 0
  di, dj = 0, 1

  rotated = 0
  while rotated < 2
    spiral[i][j] = 1

    if can_move(spiral, i, j, di, dj)
      i += di
      j += dj
      rotated = 0
    else
      di, dj = dj, -di
      rotated += 1
    end
  end   

  di, dj = -dj, di
  if spiral[i + di][j + dj] == 1
      spiral[i][j] = 0
  end
      
  spiral
end
