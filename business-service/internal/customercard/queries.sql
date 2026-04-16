-- task: Manager and cashier can add new customer card data.
-- name: CreateCustomerCard :one
INSERT INTO customer_cards (
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent;

-- task: Manager and cashier can edit customer card data.
-- name: UpdateCustomerCard :one
UPDATE customer_cards
SET cust_surname = $2,
    cust_name = $3,
    cust_patronymic = $4,
    phone_number = $5,
    city = $6,
    street = $7,
    zip_code = $8,
    percent = $9
WHERE card_number = $1
RETURNING
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent;

-- task: Manager can delete customer card data.
-- name: DeleteCustomerCard :exec
DELETE FROM customer_cards
WHERE card_number = $1;

-- task: Manager and cashier can get all customer cards sorted by customer surname.
-- name: GetAllCustomerCards :many
SELECT
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent
FROM customer_cards
ORDER BY cust_surname, cust_name, cust_patronymic;

-- task: Helper query for customer card CRUD/API flows to fetch one customer card by card number.
-- name: GetCustomerCardByNumber :one
SELECT
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent
FROM customer_cards
WHERE card_number = $1;

-- task: Manager can get all customer cards with a selected discount percent sorted by customer surname.
-- name: GetCustomerCardsByPercent :many
SELECT
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent
FROM customer_cards
WHERE percent = $1
ORDER BY cust_surname, cust_name, cust_patronymic;

-- task: Cashier can search customer cards by customer surname.
-- name: SearchCustomerCardsBySurname :many
SELECT
    card_number,
    cust_surname,
    cust_name,
    cust_patronymic,
    phone_number,
    city,
    street,
    zip_code,
    percent
FROM customer_cards
WHERE cust_surname ILIKE '%' || $1 || '%'
ORDER BY cust_surname, cust_name, cust_patronymic;
