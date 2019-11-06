-- Write a SELECT query which will return all prime numbers smaller than 100 in ascending order.
-- Your query should return one column named prime.

SELECT table1.a AS prime
FROM generate_series(2, 100) AS table1(a)
WHERE NOT EXISTS (
  SELECT table2.a
  FROM generate_series(2, table1.a-1) AS table2(a)
  WHERE table1.a % table2.a = 0
);
