-- name: CreateCheck :exec
-- task: Здійснювати продаж товарів (додавання чеків) (Cashier #7)
INSERT INTO checks (check_number, id_employee, card_number, print_date, sum_total, vat)
    VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetAllChecksOfTheDayByCashier :many
-- task: Переглянути список усіх чеків, що створив касир за цей день (Cashier #9)
SELECT
    check_number,
    id_employee,
    card_number,
    print_date,
    sum_total,
    vat
FROM
    checks
WHERE
    id_employee = $1
    AND DATE(print_date) = $2
ORDER BY
    print_date DESC;

-- name: GetAllChecksOfThePeriodByCashier :many
-- task: Переглянути список усіх чеків, що створив касир за певний період часу (Cashier #10)
SELECT
    check_number,
    id_employee,
    card_number,
    print_date,
    sum_total,
    vat
FROM
    checks
WHERE
    id_employee = $1
    AND print_date BETWEEN $2 AND $3
ORDER BY
    print_date DESC;

-- name: GetFullCheckDataByID :many
-- task: За номером чеку вивести усю інформацію про даний чек, в тому числі
-- інформацію про назву, к-сть та ціну товарів, придбаних в даному чеку (Cashier #11)
SELECT
    c.check_number,
    c.print_date,
    c.sum_total,
    c.vat,
    s.upc,
    s.product_number AS quantity,
    s.selling_price,
    p.product_name,
    e.empl_surname,
    e.empl_name,
    e.empl_patronymic
FROM
    checks c
    INNER JOIN employees e ON c.id_employee = e.id_employee
    INNER JOIN sales s ON c.check_number = s.check_number
    INNER JOIN store_products sp ON s.upc = sp.upc
    INNER JOIN products p ON sp.id_product = p.id_product
WHERE
    c.check_number = $1;

-- name: DeleteCheckByNumber :exec
-- task: Видаляти дані про чеки (Manager #3)
DELETE FROM checks
WHERE check_number = $1;

-- name: GetCheckByNumber :one
-- task: За номером чеку вивести базову інформацію про чек (Cashier #11)
SELECT
    check_number,
    id_employee,
    card_number,
    print_date,
    sum_total,
    vat
FROM
    checks
WHERE
    check_number = $1;

-- name: GetAllChecks :many
-- task: Видруковувати звіти з інформацією про усі чеки (Manager #4)
SELECT
    check_number,
    id_employee,
    card_number,
    print_date,
    sum_total,
    vat
FROM
    checks
ORDER BY
    print_date DESC;

-- name: GetCheckDetailsOfThePeriodByCashier :many
-- task: Отримати інформацію про усі чеки, створені певним касиром за певний період (Manager #17)
SELECT
    c.check_number,
    c.print_date,
    c.sum_total,
    p.product_name,
    s.upc,
    s.product_number AS quantity,
    s.selling_price
FROM
    checks c
    INNER JOIN sales s ON c.check_number = s.check_number
    INNER JOIN store_products sp ON s.upc = sp.upc
    INNER JOIN products p ON sp.id_product = p.id_product
WHERE
    id_employee = $1
    AND print_date BETWEEN $2 AND $3
ORDER BY
    c.print_date DESC;

-- name: GetCheckDetailsOfThePeriod :many
-- task: Отримати інформацію про усі чеки, створені усіма касирами за певний період (Manager #18)
SELECT
    c.check_number,
    c.print_date,
    c.sum_total,
    p.product_name,
    s.upc,
    s.product_number AS quantity,
    s.selling_price
FROM
    checks c
    INNER JOIN sales s ON c.check_number = s.check_number
    INNER JOIN store_products sp ON s.upc = sp.upc
    INNER JOIN products p ON sp.id_product = p.id_product
WHERE
    print_date BETWEEN $1 AND $2
ORDER BY
    c.print_date DESC;

-- name: GetSumOfAllChecksOfThePeriodByCashier :one
-- task: Визначити загальну суму проданих товарів з чеків, створених певним касиром (Manager #19)
SELECT
    COALESCE(SUM(sum_total), 0)::numeric(13, 4) AS total_sum
FROM
    checks
WHERE
    id_employee = $1
    AND print_date BETWEEN $2 AND $3;

-- name: GetSumOfAllChecksOfThePeriod :one
-- task: Визначити загальну суму проданих товарів з чеків, створених усіма касирами (Manager #20)
SELECT
    COALESCE(SUM(sum_total), 0)::numeric(13, 4) AS total_sum
FROM
    checks
WHERE
    print_date BETWEEN $1 AND $2;
