=begin

Write a method that takes a field for well-known board game "Battleship" as an argument and returns true if it has a valid disposition of ships, false otherwise.
Argument is guaranteed to be 10*10 two-dimension array.
Elements in the array are numbers, 0 if the cell is free and 1 if occupied by ship.

Battleship (also Battleships or Sea Battle) is a guessing game for two players. Each player has a 10x10 grid containing several "ships" and objective is to destroy enemy's forces by targetting individual cells on his field.
The ship occupies one or more cells in the grid. Size and number of ships may differ from version to version. In this kata we will use Soviet/Russian version of the game.

Before the game begins, players set up the board and place the ships accordingly to the following rules:
There must be a single battleship (size of 4 cells), 2 cruisers (size 3), 3 destroyers (size 2) and 4 submarines (size 1). Any additional ships are not allowed, as well as missing ships.
Each ship must be a straight line, except for submarines, which are just single cell.

The ship cannot overlap or be in contact with any other ship, neither by the edge nor by the corner.

=end

def validate_battlefield(field)
  visited = Array.new(10) { Array.new(10, false) }
  ships = []

  (0..9).each do |row|
    (0..9).each do |col|
      if field[row][col] == 1 && !visited[row][col]
        ship = find_ship(field, visited, row, col)
        ships << ship
      end
    end
  end

  ship_sizes = ships.map(&:length).sort.reverse
  expected_sizes = [4, 3, 3, 2, 2, 2, 1, 1, 1, 1]
  return false if ship_sizes != expected_sizes

  ships.each do |ship|
    return false unless ship_not_touching_others?(field, ship)
  end

  true
end

private

def find_ship(field, visited, start_row, start_col)
  ship_cells = []
  # check a ship type to ignore broken ships
  ship_type = nil
  horizontal_deltas = [[0, 1], [0, -1]]
  vertical_deltas = [[1, 0], [-1, 0]]
  queue = [[start_row, start_col]]

  until queue.empty?
    row, col = queue.shift
    next if visited[row][col]

    visited[row][col] = true
    ship_cells << [row, col]

    [[0, 1], [0, -1], [1, 0], [-1, 0]].each do |delta|
      if ship_type == :horizontal && !horizontal_deltas.include?(delta)
        next
      elsif ship_type == :vertical && !vertical_deltas.include?(delta)
        next
      end

      new_row, new_col = row + delta[0], col + delta[1]
      if new_row >= 0 && new_row < 10 && new_col >= 0 && new_col < 10 &&
         field[new_row][new_col] == 1 && !visited[new_row][new_col]
        ship_type = (delta == [0, 1] || delta == [0, -1]) ? :horizontal : :vertical if ship_type.nil?
        queue << [new_row, new_col]
      end
    end
  end

  ship_cells
end

def ship_not_touching_others?(field, ship)
  ship.each do |row, col|
    (-1..1).each do |delta_row|
      (-1..1).each do |delta_col|
        next if delta_row == 0 && delta_col == 0

        new_row, new_col = row + delta_row, col + delta_col
        next if new_row < 0 || new_row >= 10 || new_col < 0 || new_col >= 10

        if field[new_row][new_col] == 1 && !ship.include?([new_row, new_col])
          return false
        end
      end
    end
  end

  true
end
