-- Your task is to add up letters to one letter.

-- In SQL, you will be given a table letters, with a string column letter. Return the sum of the letters in a column letter.

-- Notes:
-- Letters will always be lowercase.
-- Letters can overflow (see second to last example of the description)
-- If no letters are given, the function should return 'z'

WITH sum_chars AS (
  SELECT SUM(COALESCE((ASCII(letter) - 96), 0))::int AS l
  FROM letters
)
SELECT CHR(
  COALESCE(NULLIF(mod(l, 26), 0), 26) + 96
) AS letter FROM sum_chars
