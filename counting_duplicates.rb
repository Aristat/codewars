def duplicate_count(text)
  text.downcase.chars.group_by{ |e| e }.select { |k, v| v.size > 1 }.map(&:first).length
end
