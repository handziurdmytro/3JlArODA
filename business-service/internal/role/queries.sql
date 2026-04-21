-- name: GetEmployeeRole :one
SELECT empl_role
FROM employees
WHERE id_employee = $1;

