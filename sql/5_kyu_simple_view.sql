-- For this challenge you need to create a VIEW. This VIEW is used by a sales store to give out vouches to members who have spent over $1000 in departments 
-- that have brought in more than $10000 total ordered by the members id. The VIEW must be called members_approved_for_voucher then you must create a SELECT query using the view.

-- resultant table schema

-- id
-- name
-- email
-- total_spending

CREATE OR REPLACE VIEW members_approved_for_voucher AS
  WITH total_sales AS (
    SELECT s.department_id, SUM(p.price) 
    FROM sales s
    JOIN products p ON s.product_id = p.id
    GROUP BY s.department_id
    HAVING SUM(p.price) > 10000
  )

  SELECT s.member_id AS id, m.name, m.email, SUM(p.price) AS total_spending 
  FROM sales s
  JOIN products p ON s.product_id = p.id
  JOIN members m ON s.member_id = m.id
  WHERE s.department_id IN (SELECT department_id FROM total_sales)
  GROUP BY (s.member_id, m.name, m.email)
  HAVING SUM(p.price) > 1000
  ORDER BY s.member_id ASC;
  
SELECT * FROM members_approved_for_voucher
