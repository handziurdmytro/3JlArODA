-- name: CreateSale :exec
-- task: Додавати товари до чеку (Cashier #7) (transaction)
INSERT INTO sales (upc, check_number, product_number, selling_price)
    VALUES ($1, $2, $3, $4);

-- name: GetProductSoldQuantity :one
-- task: Визначити загальну кількість одиниць певного товару, проданого за певний період (Manager #21)
SELECT
    COALESCE(SUM(s.product_number), 0)::bigint AS total_quantity
FROM
    sales s
    INNER JOIN checks c ON s.check_number = c.check_number
    INNER JOIN store_products sp ON s.upc = sp.upc
WHERE
    sp.id_product = $1
    AND c.print_date BETWEEN $2 AND $3;

