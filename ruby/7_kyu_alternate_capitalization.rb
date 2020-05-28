def capitalize(s)
  even_result = ""
  odd_result = ""
 
  s.chars.each_with_index do |i, index|
    if index.even?
      even_result << i.upcase
      odd_result << i
    else
      even_result << i
      odd_result << i.upcase
    end
  end
  
  [even_result, odd_result]
end
