-- task: Manager can add new product data.
-- name: CreateProduct :one
INSERT INTO products (category_number, product_name, producer, characteristics)
VALUES ($1, $2, $3, $4)
RETURNING id_product, category_number, product_name, producer, characteristics;

-- task: Manager can edit product data.
-- name: UpdateProduct :one
UPDATE products
SET category_number = $2,
    product_name = $3,
    producer = $4,
    characteristics = $5
WHERE id_product = $1
RETURNING id_product, category_number, product_name, producer, characteristics;

-- task: Manager can delete product data.
-- name: DeleteProduct :exec
DELETE FROM products
WHERE id_product = $1;

-- task: Manager and cashier can get all products sorted by product name.
-- name: GetAllProductsSortedByName :many
SELECT
    id_product,
    category_number,
    product_name,
    producer,
    characteristics
FROM products
ORDER BY product_name;

-- task: Helper query for product CRUD/API flows to fetch one product by identifier.
-- name: GetProductById :one
SELECT
    id_product,
    category_number,
    product_name,
    producer,
    characteristics
FROM products
WHERE id_product = $1;

-- task: Manager and cashier can search all products from a selected category sorted by product name.
-- name: GetProductsByCategory :many
SELECT
    id_product,
    category_number,
    product_name,
    producer,
    characteristics
FROM products
WHERE category_number = $1
ORDER BY product_name;

-- task: Cashier can search products by product name.
-- name: SearchProductsByName :many
SELECT
    id_product,
    category_number,
    product_name,
    producer,
    characteristics
FROM products
WHERE product_name ILIKE '%' || $1 || '%'
ORDER BY product_name;
