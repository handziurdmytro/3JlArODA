DROP INDEX IF EXISTS idx_sales_check_number;
DROP INDEX IF EXISTS idx_checks_print_date;
DROP INDEX IF EXISTS idx_checks_card_number;
DROP INDEX IF EXISTS idx_checks_id_employee;
DROP INDEX IF EXISTS idx_store_products_id_product;
DROP INDEX IF EXISTS idx_products_category_number;

DROP TABLE IF EXISTS sales;
DROP TABLE IF EXISTS checks;
DROP TABLE IF EXISTS store_products;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS customer_cards;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS employees;
