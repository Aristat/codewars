=begin

Create an endless stream of prime numbers - a bit like IntStream.of(2, 3, 5, 7, 11, 13, 17), but infinite. The stream must be able to produce a million primes in a few seconds.

If this is too easy, try Prime Streaming (NC-17).

Note: the require keyword is disabled.

=end

class Primes
  def self.stream
    PrimeStream.new(primes)
  end
  
  def self.primes
    @primes ||= eratosthenes(15486042)
  end
  
  private

  def self.eratosthenes(n)
    nums = [nil, nil, *2..n]
    (2..Math.sqrt(n)).each do |i|
      (i**2..n).step(i){|m| nums[m] = nil}  if nums[i]
    end
    nums.compact
  end
end

class PrimeStream
  attr_reader :primes, :index

  def initialize(primes)
    @primes = primes
    @index = -1
  end

  def next
    @index += 1
    primes[index]    
  end
end
