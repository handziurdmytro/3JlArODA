-- name: CreateCategory :one
-- task: Додавати нові дані про категорії товарів (Manager #1)
INSERT INTO categories (category_name)
VALUES ($1)
RETURNING category_number, category_name;

-- name: UpdateCategory :one
-- task: Редагувати дані про категорії товарів (Manager #2)
UPDATE categories
SET category_name = $2
WHERE category_number = $1
RETURNING category_number, category_name;

-- name: DeleteCategory :exec
-- task: Видаляти дані про категорії товарів (Manager #3)
-- todo: handle On Delete No Action (products reference categories)
DELETE FROM categories
WHERE category_number = $1;

-- name: GetAllCategoriesSortedByName :many
-- task: Отримати інформацію про усі категорії, відсортовані за назвою (Manager #8)
SELECT
    category_number,
    category_name
FROM categories
ORDER BY category_name;

-- name: GetCategoryByID :one
-- task: Helper query for category CRUD/API flows to fetch one category by identifier.
SELECT
    category_number,
    category_name
FROM categories
WHERE category_number = $1;

-- name: GetCategoryStockSummary :many
-- task: 1.	Вивести к-сть одиниць та середню ціну товарів у магазинах по кожній категорії товарів (Manager)
SELECT
    c.category_number,
    c.category_name,
    SUM(sp.products_number)                               AS total_quantity,
    ROUND(AVG(sp.selling_price), 2)                       AS avg_price
FROM categories c
    INNER JOIN products p ON c.category_number = p.category_number
    INNER JOIN store_products sp ON p.id_product = sp.id_product
GROUP BY c.category_number, c.category_name
ORDER BY c.category_name;
