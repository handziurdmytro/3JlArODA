-- name: CreateEmployee :exec
-- task: Додавати нові дані про працівників (Manager #1)
INSERT INTO employees (id_employee, empl_surname, empl_name, empl_patronymic,
                       empl_role, salary, date_of_birth, date_of_start,
                       phone_number, city, street, zip_code)
VALUES ($1, $2, $3, $4,
        $5, $6, $7, $8,
        $9, $10, $11, $12);

-- name: UpdateEmployeeByID :exec
-- task: Редагувати дані про працівників (Manager #2)
UPDATE employees
SET empl_surname    = $2,
    empl_name       = $3,
    empl_patronymic = $4,
    empl_role       = $5,
    salary          = $6,
    date_of_birth   = $7,
    date_of_start   = $8,
    phone_number    = $9,
    city            = $10,
    street          = $11,
    zip_code        = $12
WHERE id_employee = $1;

-- name: DeleteEmployeeByID :exec
-- task: Видаляти дані про працівників (Manager #3)
-- todo: handle On Delete No Action
DELETE
FROM employees
WHERE id_employee = $1;

-- name: GetAllEmployees :many
-- task: Отримати інформацію про усіх працівників, відсортованих за прізвищем (Manager #5)
SELECT id_employee,
       empl_surname,
       empl_name,
       empl_patronymic,
       empl_role,
       salary,
       date_of_birth,
       date_of_start,
       phone_number,
       city,
       street,
       zip_code
FROM employees
ORDER BY empl_surname, empl_name, empl_patronymic;

-- name: GetEmployeeByID :one
-- task: Можливість отримати усю інформацію про себе. (Cashier #15)
SELECT id_employee,
       empl_surname,
       empl_name,
       empl_patronymic,
       empl_role,
       salary,
       date_of_birth,
       date_of_start,
       phone_number,
       city,
       street,
       zip_code
FROM employees
WHERE id_employee = $1;

-- name: GetEmployeesByRole :many
-- task: Отримати інформацію про усіх працівників, що займають посаду касира,
-- відсортованих за прізвищем (Manager #6)
SELECT id_employee,
       empl_surname,
       empl_name,
       empl_patronymic,
       empl_role,
       salary,
       date_of_birth,
       date_of_start,
       phone_number,
       city,
       street,
       zip_code
FROM employees
WHERE empl_role = $1
ORDER BY empl_surname, empl_name, empl_patronymic;

-- name: GetEmployeeDataBySurname :many
-- task: За прізвищем працівника знайти його телефон та адресу (Manager #11)
SELECT phone_number,
       city,
       street,
       zip_code
FROM employees
WHERE empl_surname = $1;

-- name: GetEmployeeDataByFullName :many
-- task: За прізвищем працівника знайти його телефон та адресу (Manager #11)
SELECT phone_number,
       city,
       street,
       zip_code
FROM employees
WHERE empl_surname = $1
  AND empl_name = $2
  AND empl_patronymic = $3;
