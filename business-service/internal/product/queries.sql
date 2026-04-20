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

-- task: Manager can analyze sold quantity and revenue for products from a selected category during a selected period.
-- name: GetProductSalesStatsByCategoryAndPeriod :many
SELECT
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name,
    COALESCE(SUM(s.product_number), 0)::bigint AS total_sold_quantity,
    COALESCE(SUM(s.product_number * s.selling_price), 0)::numeric(13, 4) AS total_revenue
FROM products p
INNER JOIN categories c ON p.category_number = c.category_number
INNER JOIN store_products sp ON p.id_product = sp.id_product
INNER JOIN sales s ON sp.upc = s.upc
INNER JOIN checks ch ON s.check_number = ch.check_number
WHERE
    c.category_number = $1
    AND ch.print_date >= $2
    AND ch.print_date <= $3
GROUP BY
    p.id_product,
    p.product_name,
    p.producer,
    p.characteristics,
    c.category_number,
    c.category_name
ORDER BY
    total_sold_quantity DESC,
    p.product_name;
