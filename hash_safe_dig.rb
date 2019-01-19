=begin

Implement a safe version of Ruby's

Given a hash and a list of n keys, return the value of the nth nested value.

If there are more keys than there are nested levels of hashes, the method should return nil rather than raise an exception.

All keys are assumed to be symbols in this kata.

=end


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
