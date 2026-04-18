# API Endpoints

Base URL prefix:

```text
/api/v1
```

All endpoints except `/auth/register` and `/auth/login` are protected and require:

```http
Authorization: Bearer <token>
```

Date/time fields are JSON strings. Use RFC3339 unless the handler/service later documents a stricter format, for example:

```json
"2026-04-18T12:00:00Z"
```

## Auth

### POST `/api/v1/auth/register`

Request body:

```json
{
  "username": "manager",
  "password": "password123"
}
```

### POST `/api/v1/auth/login`

Request body:

```json
{
  "username": "manager",
  "password": "password123"
}
```

## Employee

### GET `/api/v1/employee/`

No request body.

### GET `/api/v1/employee/:id`

Path params:

```text
id: string
```

No request body.

### POST `/api/v1/employee/`

Request body:

```json
{
  "id": "EMP001",
  "surname": "Ivanenko",
  "name": "Ivan",
  "patronymic": "Ivanovych",
  "role": "cashier",
  "salary": 15000.0,
  "date_of_birth": "1999-01-15T00:00:00Z",
  "date_of_start": "2024-09-01T00:00:00Z",
  "phone_number": "+380501112233",
  "city": "Kyiv",
  "street": "Khreshchatyk 1",
  "zip_code": "01001"
}
```

`patronymic` may be `null`.

### PUT `/api/v1/employee/:id`

Path params:

```text
id: string
```

Request body:

```json
{
  "surname": "Ivanenko",
  "name": "Ivan",
  "patronymic": "Ivanovych",
  "role": "cashier",
  "salary": 15000.0,
  "date_of_birth": "1999-01-15T00:00:00Z",
  "date_of_start": "2024-09-01T00:00:00Z",
  "phone_number": "+380501112233",
  "city": "Kyiv",
  "street": "Khreshchatyk 1",
  "zip_code": "01001"
}
```

`PUT` is a full update. Send all fields. `patronymic` may be `null`.

### DELETE `/api/v1/employee/:id`

Path params:

```text
id: string
```

No request body.

## Category

### GET `/api/v1/category/`

No request body.

### GET `/api/v1/category/:number`

Path params:

```text
number: integer
```

No request body.

### POST `/api/v1/category/`

Request body:

```json
{
  "name": "Dairy"
}
```

### PUT `/api/v1/category/:number`

Path params:

```text
number: integer
```

Request body:

```json
{
  "name": "Dairy"
}
```

`PUT` is a full update. Send all fields.

### DELETE `/api/v1/category/:number`

Path params:

```text
number: integer
```

No request body.

## Product

### GET `/api/v1/product/`

No request body.

### GET `/api/v1/product/:id`

Path params:

```text
id: integer
```

No request body.

### POST `/api/v1/product/`

Request body:

```json
{
  "category_number": 1,
  "name": "Milk",
  "producer": "Farm Ltd",
  "characteristics": "2.5% fat, 1L"
}
```

`producer` may be `null`.

### PUT `/api/v1/product/:id`

Path params:

```text
id: integer
```

Request body:

```json
{
  "category_number": 1,
  "name": "Milk",
  "producer": "Farm Ltd",
  "characteristics": "2.5% fat, 1L"
}
```

`PUT` is a full update. Send all fields. `producer` may be `null`.

### DELETE `/api/v1/product/:id`

Path params:

```text
id: integer
```

No request body.

## Store Product

### GET `/api/v1/store-product/`

No request body.

### GET `/api/v1/store-product/:upc`

Path params:

```text
upc: string
```

No request body.

### POST `/api/v1/store-product/`

Request body:

```json
{
  "upc": "123456789012",
  "upc_prom": null,
  "product_id": 1,
  "selling_price": 42.5,
  "products_number": 100,
  "promotional_product": false
}
```

`upc_prom` may be `null`.

### PUT `/api/v1/store-product/:upc`

Path params:

```text
upc: string
```

Request body:

```json
{
  "upc_prom": null,
  "product_id": 1,
  "selling_price": 42.5,
  "products_number": 100,
  "promotional_product": false
}
```

`PUT` is a full update. Send all fields. `upc_prom` may be `null`.

### DELETE `/api/v1/store-product/:upc`

Path params:

```text
upc: string
```

No request body.

## Customer Card

### GET `/api/v1/customer-card/`

No request body.

### GET `/api/v1/customer-card/:number`

Path params:

```text
number: string
```

No request body.

### POST `/api/v1/customer-card/`

Request body:

```json
{
  "card_number": "CARD000000001",
  "surname": "Shevchenko",
  "name": "Taras",
  "patronymic": "Hryhorovych",
  "phone_number": "+380501112233",
  "city": "Kyiv",
  "street": "Volodymyrska 1",
  "zip_code": "01001",
  "percent": 5
}
```

`patronymic`, `city`, `street`, and `zip_code` may be `null`.

### PUT `/api/v1/customer-card/:number`

Path params:

```text
number: string
```

Request body:

```json
{
  "surname": "Shevchenko",
  "name": "Taras",
  "patronymic": "Hryhorovych",
  "phone_number": "+380501112233",
  "city": "Kyiv",
  "street": "Volodymyrska 1",
  "zip_code": "01001",
  "percent": 5
}
```

`PUT` is a full update. Send all fields. `patronymic`, `city`, `street`, and `zip_code` may be `null`.

### DELETE `/api/v1/customer-card/:number`

Path params:

```text
number: string
```

No request body.

## Check

### GET `/api/v1/check/`

No request body.

### GET `/api/v1/check/:number`

Path params:

```text
number: string
```

No request body.

### POST `/api/v1/check/`

Request body:

```json
{
  "number": "CHK000000001",
  "employee_id": "EMP001",
  "card_number": "CARD000000001",
  "print_date": "2026-04-18T12:00:00Z",
  "sum_total": 120.5,
  "vat": 20.08
}
```

`card_number` may be `null`.

### DELETE `/api/v1/check/:number`

Path params:

```text
number: string
```

No request body.

## Sale

### GET `/api/v1/sale/`

No request body.

### GET `/api/v1/sale/item`

No request body.

### POST `/api/v1/sale/`

Request body:

```json
{
  "upc": "123456789012",
  "check_number": "CHK000000001",
  "product_number": 2,
  "selling_price": 42.5
}
```
