ARRAY = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('').freeze

def dec2FactString(nb)
    f = 1
    n = 2

    while f < nb
      if n * f < nb
        f *= n
        n += 1
      else
        n -= 1
        break
      end
    end

    res = ''
    while n > 0
      res += ARRAY[nb / f]
      nb %= f
      f /= n
      n -= 1
    end

    res + '0'
end

def factString2Dec(string)
    s = 0
    f = 1
    string = string.reverse.split('')

    string.each_with_index do |i, index|
      next if index == 0
      f *= index
      s += f * ARRAY.find_index(i)
    end

    s
end
