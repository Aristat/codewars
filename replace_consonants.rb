CONSONANTS = ["B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "R", "S", "T", "V", "W", "X", "Z", "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "x", "z"]

def replace(string)
  new_array = []
  last_index = nil
  array_string = string.split('')

  array_string.each_with_index do |v, i|
    unless CONSONANTS.include?(v)
      new_array << v 
      next
    end
    

    array_string.reverse.each_with_index do |vv, ii|
      next if !last_index.nil? && ii <= last_index

      if CONSONANTS.include?(vv)
        new_array << vv
        last_index = ii

        break
      end
    end
  end

  new_array.join('')
end
