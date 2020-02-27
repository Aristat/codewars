For this challenge you need to create a SELECT statement, this SELECT statement will use an EXISTS to check whether a department has had a sale with a price over 98.00 dollars.

departments table schema
id
name
sales table schema
id
department_id (department foreign key)
name
price
card_name
card_number
transaction_date
resultant table schema
id
name
NOTE: Your solu

SELECT departments.id, departments.name
FROM departments
WHERE EXISTS (SELECT 1 FROM sales WHERE sales.department_id = departments.id AND sales.price > 98)
