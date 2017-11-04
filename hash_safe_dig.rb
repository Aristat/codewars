class Hash
  def safe_dig(*keys)
    nested_hash = self
    keys_length = keys.length
    
    keys.each_with_index do |key, index|
      nested_hash = nested_hash[key]
      
      if !nested_hash.kind_of?(Hash)
        nested_hash = nil if keys_length != (index + 1)
        break 
      end 
    end
    
    nested_hash
  end
end
