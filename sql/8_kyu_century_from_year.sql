-- Introduction
-- The first century spans from the year 1 up to and including the year 100, The second - from the year 101 up to and including the year 200, etc.

SELECT EXTRACT(CENTURY FROM (SELECT to_date(yr::text, 'YYYY'))) AS century FROM years
