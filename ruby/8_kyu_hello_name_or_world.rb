=begin

Define a method hello that returns "Hello, Name!" to a given name, or says Hello, World! if name is not given (or passed as an empty String).

Assuming that name is a String and it checks for user typos to return a name with a first capital letter (Xxxx).

=end

def hello(name = nil)
  (name.nil? || name == '') ? "Hello, World!" : "Hello, #{name.capitalize}!"
end
