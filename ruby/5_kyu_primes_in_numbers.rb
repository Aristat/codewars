def primeFactors(n)
  result = {}
  n, result = prime_decomposition(n, 2, result)
  n, result = prime_decomposition(n, 3, result)

  prime = 5

  while prime * prime <= n
    n, result = prime_decomposition(n, prime, result)
    prime += 2
    n, result = prime_decomposition(n, prime, result)
    prime += 4
  end

  result[n] = 1 if n > 1

  s = ""
  result.each do |k, v|
    if v == 1
      s += "(#{k})"
    else
      s += "(#{k}**#{v})"
    end
  end
  s
end


def prime_decomposition(n, prime, result)
  factors = 0

  while(q, r = n.divmod(prime)
        r.zero?)
    factors += 1
    n = q
  end

  if factors != 0
    result[prime] = factors
  end

  [n, result]
end
