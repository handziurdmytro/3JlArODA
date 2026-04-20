-- name: CreateStoreProduct :one
-- task: Додавати нові дані про товари у магазині (Manager #1)
INSERT INTO store_products (upc, upc_prom, id_product, selling_price, products_number, promotional_product)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    upc,
    upc_prom,
    id_product,
    selling_price,
    products_number,
    promotional_product;

-- name: UpdateStoreProduct :one
-- task: Редагувати дані про товари у магазині (Manager #2)
UPDATE store_products
SET
    upc_prom            = $2,
    id_product          = $3,
    selling_price       = $4,
    products_number     = $5,
    promotional_product = $6
WHERE upc = $1
RETURNING
    upc,
    upc_prom,
    id_product,
    selling_price,
    products_number,
    promotional_product;

-- name: DeleteStoreProduct :exec
-- task: Видаляти дані про товари у магазині (Manager #3)
-- todo: handle On Delete No Action (sales reference store_products)
DELETE FROM store_products
WHERE upc = $1;

-- name: GetStoreProductByUPC :one
-- task: За UPC знайти ціну продажу, к-сть, назву та характеристики товару (Manager #14, Cashier #14)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE sp.upc = $1;

-- name: GetAllStoreProductsSortedByQuantity :many
-- task: Отримати інформацію про усі товари у магазині, відсортовані за кількістю (Manager #10)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
ORDER BY sp.products_number DESC;

-- name: GetAllStoreProductsSortedByName :many
-- task: Отримати інформацію про усі товари у магазині, відсортовані за назвою (Cashier #2)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
ORDER BY p.product_name;

-- name: GetPromoStoreProductsSortedByQuantity :many
-- task: Отримати інформацію про усі акційні товари, відсортовані за кількістю (Manager #15, Cashier #12)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE sp.promotional_product = TRUE
ORDER BY sp.products_number DESC;

-- name: GetPromoStoreProductsSortedByName :many
-- task: Отримати інформацію про усі акційні товари, відсортовані за назвою (Manager #15, Cashier #12)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE sp.promotional_product = TRUE
ORDER BY p.product_name;

-- name: GetNonPromoStoreProductsSortedByQuantity :many
-- task: Отримати інформацію про усі не акційні товари, відсортовані за кількістю (Manager #16, Cashier #13)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE sp.promotional_product = FALSE
ORDER BY sp.products_number DESC;

-- name: GetNonPromoStoreProductsSortedByName :many
-- task: Отримати інформацію про усі не акційні товари, відсортовані за назвою (Manager #16, Cashier #13)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE sp.promotional_product = FALSE
ORDER BY p.product_name;

-- name: GetStoreProductsByCategorySortedByName :many
-- task: Здійснити пошук товарів, що належать певній категорії (Manager #13, Cashier #5)
SELECT
    sp.upc,
    sp.upc_prom,
    sp.selling_price,
    sp.products_number,
    sp.promotional_product,
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
FROM store_products sp
    INNER JOIN products p ON sp.id_product = p.id_product
    INNER JOIN categories c ON p.category_number = c.category_number
WHERE c.category_number = $1
ORDER BY p.product_name;

-- name: GetCashiersWhoSoldAllProductsFromCategory :many
-- task: Знайти касирів, які продали кожен товар з певної категорії за період (Manager)
-- $1 — category_number, $2 — date_from, $3 — date_to
SELECT
    e.id_employee,
    e.empl_surname,
    e.empl_name,
    e.empl_patronymic,
    e.phone_number
FROM employees e
WHERE
    e.empl_role = 'cashier'
    AND NOT EXISTS (
        SELECT 1
        FROM products p
        WHERE p.category_number = $1
        AND NOT EXISTS (
            SELECT 1
            FROM checks c
                INNER JOIN sales s ON c.check_number = s.check_number
                INNER JOIN store_products sp ON s.upc = sp.upc
            WHERE
                c.id_employee  = e.id_employee
                AND sp.id_product = p.id_product
                AND c.print_date BETWEEN $2 AND $3
        )
    )
ORDER BY e.empl_surname, e.empl_name, e.empl_patronymic;