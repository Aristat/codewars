=begin
Sudoku Solution Validator

Sudoku is a game played on a 9x9 grid. The goal of the game is to fill all cells of the grid with digits from 1 to 9, so that each column, each row, and each of the nine 3x3 sub-grids (also known as blocks) contain all of the digits from 1 to 9. 
(More info at: http://en.wikipedia.org/wiki/Sudoku)

Sudoku Solution Validator

Write a function validSolution/ValidateSolution/valid_solution() that accepts a 2D array representing a Sudoku board, and returns true if it is a valid solution, or false otherwise. The cells of the sudoku board may also contain 0's, which will represent empty cells. Boards containing one or more zeroes are considered to be invalid solutions.

The board is always 9 cells by 9 cells, and every cell only contains integers from 0 to 9
=end

def valid_solution(board)
  transpose_board = board.transpose

  sub_grid_valid = true
  sub_grids = [[0,1,2], [3,4,5], [6,7,8]]
  
  sub_grids.each do |grid_x|
    sub_grids.each do |grid_y|
      array = []
      
      grid_y.each do |y|
        array = array | board[y][grid_x.first, 3]
      end
      
      if array.sum != 45
        sub_grid_valid = false
        break
      end
    end
  end
  
  sub_grid_valid && board.all? { |row| row.uniq.sum == 45 } && transpose_board.all? { |row| row.uniq.sum == 45 }
end
