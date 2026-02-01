=begin
In a grid of 4 by 4 squares you want to place a skyscraper in each square with only some clues:

The height of the skyscrapers is between 1 and 4
No two skyscrapers in a row or column may have the same number of floors
A clue is the number of skyscrapers that you can see in a row or column from the outside
Higher skyscrapers block the view of lower skyscrapers located behind them

Can you write a program that can solve this puzzle?

Task:

Finish:
def solve_puzzle(clues)

The clues are passed in as an array of 16 items. This array contains the clues around the clock, indexed as follows:

  	 0	 1	   2	   3
 15	  	  	  	  	 4
 14	  	  	  	  	 5
 13	  	  	  	  	 6
 12	  	  	  	  	 7
  	11	10	 9	 8

If no clue is available, the value `0` is used
Each puzzle has only one possible solution
`SolvePuzzle()` should return a `int[][]` matrix. The first index is the row, the second index is the column. (Python: return a 4-tuple of 4-tuples, Ruby: 4-Array of 4-Arrays)

=end

def solve_puzzle(clues)
  grid = Array.new(4) { Array.new(4, 0) }
  solve_grid(grid, 0, 0, clues) ? grid : nil
end

def solve_grid(grid, row, col, clues)
  return is_valid_solution(grid, clues) if row == 4

  next_row, next_col =
    if col == 3
      [row + 1, 0]
    else
      [row, col + 1]
    end

  # Try each height (1-4) for the current cell
  (1..4).each do |height|
    if is_valid_placement(grid, row, col, height)
      grid[row][col] = height

      # Recursively try to fill the rest
      if solve_grid(grid, next_row, next_col, clues)
        return true
      end

      grid[row][col] = 0
    end
  end

  false
end

# height already exists or not
def is_valid_placement(grid, row, col, height)
  (0...col).each do |c|
    return false if grid[row][c] == height
  end

  (0...row).each do |r|
    return false if grid[r][col] == height
  end

  true
end

# Check all clues
def is_valid_solution(grid, clues)
  16.times do |i|
    next if clues[i] == 0

    visible_count =
      case i
      when 0..3 # Top
        visible = 0
        max_height = 0

        4.times do |row|
          if grid[row][i] > max_height
            visible += 1
            max_height = grid[row][i]
          end
        end

        visible
      when 4..7 # Right
        visible = 0
        max_height = 0

        3.downto(0) do |col|
          if grid[i - 4][col] > max_height
            visible += 1
            max_height = grid[i - 4][col]
          end
        end

        visible
      when 8..11 # Bottom
        visible = 0
        max_height = 0

        3.downto(0) do |row|
          if grid[row][11 - i] > max_height
            visible += 1
            max_height = grid[row][11 - i]
          end
        end

        visible
      when 12..15 # Left
        visible = 0
        max_height = 0

        4.times do |col|
          if grid[15 - i][col] > max_height
            visible += 1
            max_height = grid[15 - i][col]
          end
        end

        visible
      end

    return false if visible_count != clues[i]
  end

  true
end
