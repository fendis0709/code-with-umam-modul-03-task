# Category, Product & Transaction API

A REST API for managing categories, products, and transactions built with Go and PostgreSQL, following clean architecture principles.

## Course Reference

This project is part of the "Bootcamp Jago Golang - Code With Umam" course on [CodeWithUmam - Course Online](https://docs.kodingworks.io/s/820d006c-a994-4487-b993-bc3b4171a35d) and [CodeWithUmam - Youtube](https://www.youtube.com/watch?v=5rJ5g8knuRU).

This repo is submission for Modul 03 Task on Week 03 [CodeWithUmam - Course Task #03](https://docs.kodingworks.io/s/820d006c-a994-4487-b993-bc3b4171a35d#h-task-session-3).

## Architecture

This project follows clean architecture principles with the following layers:
- **Handler**: HTTP request/response handling
- **Service**: Business logic layer
- **Repository**: Data access layer
- **Model**: Domain entities
- **Transport**: Request/response DTOs
- **Database**: PostgreSQL database connection
- **Config**: Application configuration

## How to Use Locally

### Prerequisites
- Go 1.25 or higher installed on your system
- PostgreSQL database

### Environment Configuration

Create a `.env` file in the project root with the following variables:

```env
APP_PORT=6969
DB_CONN=postgres://username:password@localhost:5432/dbname?sslmode=disable
```

Or set them as environment variables:
```bash
export APP_PORT=6969
export DB_CONN="postgres://username:password@localhost:5432/dbname?sslmode=disable"
```

### Database Setup

Ensure you have PostgreSQL installed and create the necessary tables:

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    stock INTEGER,
    price DECIMAL(10, 2),
    category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    purchased_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transactions(id),
    product_id INTEGER REFERENCES products(id),
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INTEGER NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    purchased_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

### Running the Application

1. Navigate to the project directory:
```bash
cd /<your-path-project>/modul-03-task
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

4. The server will start on `http://localhost:6969`

You should see:
```
Database connected successfully.
Server is up and running
http://localhost:6969
```

## API Endpoints

### General
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | Get all categories |
| POST | `/categories` | Create a new category |
| GET | `/categories/{uuid}` | Get a specific category |
| PUT | `/categories/{uuid}` | Update a category |
| DELETE | `/categories/{uuid}` | Delete a category |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| POST | `/products` | Create a new product |
| GET | `/products/{uuid}` | Get a specific product |
| PUT | `/products/{uuid}` | Update a product |
| DELETE | `/products/{uuid}` | Delete a product |

### Checkout
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/checkouts` | Create a checkout transaction |

### Reports
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/reports` | Get report by date range (query params: start_date, end_date) |
| GET | `/reports/hari-ini` | Get today's report |

## API Usage with cURL

### 1. Health Check
Check if the server is running.

```bash
curl -X GET http://localhost:6969/
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

---

## Category Endpoints

### 2. Get All Categories
Retrieve all categories.

```bash
curl -X GET http://localhost:6969/categories
```

**Response:**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }
]
```

---

### 3. Create a New Category
Add a new category to the system.

```bash
curl -X POST http://localhost:6969/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }'
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Electronics",
  "description": "Electronic devices and accessories"
}
```

---

### 4. Get Category by UUID
Retrieve a specific category by its UUID.

```bash
curl -X GET http://localhost:6969/categories/550e8400-e29b-41d4-a716-446655440000
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Electronics",
  "description": "Electronic devices and accessories"
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 5. Update a Category
Update an existing category by its UUID.

```bash
curl -X PUT http://localhost:6969/categories/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics & Gadgets",
    "description": "All types of electronic devices and gadgets"
  }'
```

**Response:**
```
HTTP 200 OK
(empty body)
```

**Error Response (Not Found):**
```
Not Found
```

---

### 6. Delete a Category
Remove a category from the system.

```bash
curl -X DELETE http://localhost:6969/categories/550e8400-e29b-41d4-a716-446655440000
```

**Response:**
```
HTTP 200 OK
(empty body)
```

**Error Response (Not Found):**
```
Not Found
```

---

## Product Endpoints

### 7. Get All Products
Retrieve all products with their associated categories.

```bash
curl -X GET http://localhost:6969/products
```

**Response:**
```json
[
  {
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "name": "iPhone 15",
    "stock": 50,
    "price": 999.99,
    "category": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Electronics",
      "description": "Electronic devices and accessories"
    }
  }
]
```

---

### 8. Create a New Product
Add a new product to the system.

```bash
curl -X POST http://localhost:6969/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "stock": 50,
    "price": 999.99,
    "category_id": "550e8400-e29b-41d4-a716-446655440000"
  }'
```

**Response:**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "name": "iPhone 15",
  "stock": 50,
  "price": 999.99,
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }
}
```

---

### 9. Get Product by UUID
Retrieve a specific product by its UUID.

```bash
curl -X GET http://localhost:6969/products/660e8400-e29b-41d4-a716-446655440000
```

**Response:**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "name": "iPhone 15",
  "stock": 50,
  "price": 999.99,
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 10. Update a Product
Update an existing product by its UUID.

```bash
curl -X PUT http://localhost:6969/products/660e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15 Pro",
    "stock": 30,
    "price": 1199.99,
    "category_id": "550e8400-e29b-41d4-a716-446655440000"
  }'
```

**Response:**
```
HTTP 200 OK
(empty body)
```

**Error Response (Not Found):**
```
Not Found
```

---

### 11. Delete a Product
Remove a product from the system.

```bash
curl -X DELETE http://localhost:6969/products/660e8400-e29b-41d4-a716-446655440000
```

**Response:**
```
HTTP 200 OK
(empty body)
```

**Error Response (Not Found):**
```
Not Found
```

---

## Checkout Endpoints

### 12. Create a Checkout Transaction
Create a new checkout transaction with multiple products.

```bash
curl -X POST http://localhost:6969/checkouts \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440000",
        "quantity": 2
      },
      {
        "id": "770e8400-e29b-41d4-a716-446655440001",
        "quantity": 1
      }
    ]
  }'
```

**Response:**
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440000",
  "date": "2026-02-08T10:30:00Z",
  "total_amount": 2999.97,
  "items": [
    {
      "product_id": "660e8400-e29b-41d4-a716-446655440000",
      "product_name": "iPhone 15",
      "quantity": 2,
      "unit_price": 999.99,
      "total_price": 1999.98
    },
    {
      "product_id": "770e8400-e29b-41d4-a716-446655440001",
      "product_name": "MacBook Pro",
      "quantity": 1,
      "unit_price": 999.99,
      "total_price": 999.99
    }
  ]
}
```

**Error Response (No Products Found):**
```
No Products Found
```

---

## Report Endpoints

### 13. Get Today's Report
Retrieve today's sales report including total revenue, transaction count, and most purchased item.

```bash
curl -X GET http://localhost:6969/reports/hari-ini
```

**Response:**
```json
{
  "total_revenue": 15999.95,
  "total_transaksi": 8,
  "produk_terlaris": {
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "nama": "iPhone 15",
    "qty_terjual": 25
  }
}
```

---

### 14. Get Report by Date Range
Retrieve sales report for a specific date range.

```bash
curl -X GET "http://localhost:6969/reports?start_date=2026-02-01&end_date=2026-02-08"
```

**Query Parameters:**
- `start_date`: Start date in YYYY-MM-DD format (optional)
- `end_date`: End date in YYYY-MM-DD format (optional)

**Response:**
```json
{
  "total_revenue": 45999.85,
  "total_transaksi": 23,
  "produk_terlaris": {
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "nama": "iPhone 15",
    "qty_terjual": 67
  }
}
```

---

## Data Structures

### Category Response
```json
{
  "id": "string (UUID v4, auto-generated)",
  "name": "string",
  "description": "string"
}
```

### Category Request (POST/PUT)
```json
{
  "name": "string (required)",
  "description": "string (required)"
}
```

### Product Response
```json
{
  "id": "string (UUID v4, auto-generated)",
  "name": "string",
  "stock": "integer (nullable)",
  "price": "float (nullable)",
  "category": {
    "id": "string (UUID)",
    "name": "string",
    "description": "string"
  }
}
```

### Product Request (POST/PUT)
```json
{
  "name": "string (required)",
  "stock": "integer (optional)",
  "price": "float (optional)",
  "category_id": "string (optional, category UUID)"
}
```

### Checkout Request (POST)
```json
{
  "items": [
    {
      "id": "string (required, product UUID)",
      "quantity": "integer (required)"
    }
  ]
}
```

### Checkout Response
```json
{
  "id": "string (UUID v4, auto-generated)",
  "date": "string (ISO 8601 timestamp)",
  "total_amount": "float",
  "items": [
    {
      "product_id": "string (UUID)",
      "product_name": "string",
      "quantity": "integer",
      "unit_price": "float",
      "total_price": "float"
    }
  ]
}
```

### Report Response
```json
{
  "total_revenue": "float",
  "total_transaksi": "integer",
  "produk_terlaris": {
    "id": "string (UUID)",
    "nama": "string",
    "qty_terjual": "integer"
  }
}
```

## Notes
- The application uses PostgreSQL for data persistence
- UUIDs are automatically generated using UUID v4 format for categories, products, and transactions
- The `id` field in responses is the UUID (string), not the database integer ID
- All responses are in JSON format
- Product prices are stored with 2 decimal precision
- Products can optionally be associated with a category using `category_id` (UUID string) in requests
- When fetching products, the full category details are included in the nested `category` object if associated
- Response arrays are returned directly (not wrapped in a data object)
- Checkout transactions automatically update product stock quantities
- Checkout transactions calculate total amounts based on current product prices
- Reports aggregate transaction data and identify the most purchased products
- Date range queries in reports use YYYY-MM-DD format
