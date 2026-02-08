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

### Deployed API

This API is also deployed and accessible at:
```
https://fendi-modul-03-task.up.railway.app
```

You can test the API using either:
- **Local**: `http://localhost:6969` (requires running the application locally)
- **Deployed**: `https://fendi-modul-03-task.up.railway.app` (publicly accessible)

**Note**: In all the cURL examples below, you can replace `http://localhost:6969` with `https://fendi-modul-03-task.up.railway.app` to test against the deployed version.

## API Endpoints

### General
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | Get all categories |
| GET | `/categories?search={keyword}` | Search categories by name |
| POST | `/categories` | Create a new category |
| GET | `/categories/{uuid}` | Get a specific category |
| PUT | `/categories/{uuid}` | Update a category |
| DELETE | `/categories/{uuid}` | Delete a category |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/products?search={keyword}` | Search products by name |
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

**Base URLs:**
- Local: `http://localhost:6969`
- Deployed: `https://fendi-modul-03-task.up.railway.app`

*Note: All examples use `http://localhost:6969`. Replace with the deployed URL if testing the live API.*

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
    "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
    "name": "Makanan",
    "description": null
  },
  {
    "id": "b9d3398b-5039-4c40-84fc-c8299cb5926b",
    "name": "Minuman",
    "description": null
  }
]
```

---

### 3. Search Categories
Search for categories by name.

```bash
curl -X GET "http://localhost:6969/categories?search=makan"
```

**Response:**
```json
[
  {
    "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
    "name": "Makanan",
    "description": null
  }
]
```

---

### 4. Create a New Category
Add a new category to the system.

```bash
curl -X POST http://localhost:6969/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Handcrafted Steel Ball",
    "description": "Repudiandae velit vel totam quae molestiae odit autem velit."
  }'
```

**Response:**
```json
{
  "id": "0259e3a9-22d4-4686-aaaf-1006b832aff7",
  "name": "Handcrafted Steel Ball",
  "description": "Repudiandae velit vel totam quae molestiae odit autem velit."
}
```

---

### 5. Get Category by UUID
Retrieve a specific category by its UUID.

```bash
curl -X GET http://localhost:6969/categories/b05d2319-dd1b-4151-803d-8e7de6efd9d0
```

**Response:**
```json
{
  "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
  "name": "Makanan",
  "description": null
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 6. Update a Category
Update an existing category by its UUID.

```bash
curl -X PUT http://localhost:6969/categories/0259e3a9-22d4-4686-aaaf-1006b832aff7 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Licensed Concrete Table",
    "description": "Maxime aspernatur vel et qui."
  }'
```

**Response:**
```json
{
  "id": "0259e3a9-22d4-4686-aaaf-1006b832aff7",
  "name": "Licensed Concrete Table",
  "description": "Maxime aspernatur vel et qui."
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 7. Delete a Category
Remove a category from the system.

```bash
curl -X DELETE http://localhost:6969/categories/0259e3a9-22d4-4686-aaaf-1006b832aff7
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

**Error Response (Not Found):**
```
Not Found
```

---

## Product Endpoints

### 8. Get All Products
Retrieve all products with their associated categories.

```bash
curl -X GET http://localhost:6969/products
```

**Response:**
```json
[
  {
    "id": "8a046717-8407-4b22-b019-f7af47949c83",
    "name": "Indomie Goreng",
    "stock": 100,
    "price": 2500,
    "category": {
      "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
      "name": "Makanan",
      "description": null
    }
  }
]
```

---

### 9. Search Products
Search for products by name.

```bash
curl -X GET "http://localhost:6969/products?search=indo"
```

**Response:**
```json
[
  {
    "id": "8a046717-8407-4b22-b019-f7af47949c83",
    "name": "Indomie Goreng",
    "stock": 100,
    "price": 2500,
    "category": {
      "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
      "name": "Makanan",
      "description": null
    }
  }
]
```

---

### 10. Create a New Product
Add a new product to the system.

```bash
curl -X POST http://localhost:6969/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng",
    "stock": 100,
    "price": 2500,
    "category_id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0"
  }'
```

**Response:**
```json
{
  "id": "8a046717-8407-4b22-b019-f7af47949c83",
  "name": "Indomie Goreng",
  "stock": 100,
  "price": 2500,
  "category": {
    "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
    "name": "Makanan",
    "description": null
  }
}
```

---

### 11. Get Product by UUID
Retrieve a specific product by its UUID.

```bash
curl -X GET http://localhost:6969/products/8a046717-8407-4b22-b019-f7af47949c83
```

**Response:**
```json
{
  "id": "8a046717-8407-4b22-b019-f7af47949c83",
  "name": "Indomie Goreng",
  "stock": 100,
  "price": 2500,
  "category": {
    "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
    "name": "Makanan",
    "description": null
  }
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 12. Update a Product
Update an existing product by its UUID.

```bash
curl -X PUT http://localhost:6969/products/69ad9789-e397-42ff-a551-f37e452c2a44 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Licensed Concrete Car",
    "stock": 15,
    "price": 3500,
    "category_id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0"
  }'
```

**Response:**
```json
{
  "id": "69ad9789-e397-42ff-a551-f37e452c2a44",
  "name": "Licensed Concrete Car",
  "stock": 15,
  "price": 3500,
  "category": {
    "id": "b05d2319-dd1b-4151-803d-8e7de6efd9d0",
    "name": "Makanan",
    "description": null
  }
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 13. Delete a Product
Remove a product from the system.

```bash
curl -X DELETE http://localhost:6969/products/69ad9789-e397-42ff-a551-f37e452c2a44
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

**Error Response (Not Found):**
```
Not Found
```

---

## Checkout Endpoints

### 14. Create a Checkout Transaction
Create a new checkout transaction with multiple products.

```bash
curl -X POST http://localhost:6969/checkouts \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "id": "8a046717-8407-4b22-b019-f7af47949c83",
        "quantity": 2
      }
    ]
  }'
```

**Response:**
```json
{
  "id": "9d5898fb-19d2-4878-b76f-c841679bfda4",
  "date": "2026-02-08",
  "total_amount": 5000,
  "items": [
    {
      "product_id": "8a046717-8407-4b22-b019-f7af47949c83",
      "product_name": "Indomie Goreng",
      "quantity": 2,
      "unit_price": 2500,
      "total_price": 5000
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

### 15. Get Today's Report
Retrieve today's sales report including total revenue, transaction count, and most purchased item.

```bash
curl -X GET http://localhost:6969/reports/hari-ini
```

**Response:**
```json
{
  "total_revenue": 17500,
  "total_transaksi": 6,
  "produk_terlaris": {
    "id": "8a046717-8407-4b22-b019-f7af47949c83",
    "nama": "Indomie Goreng",
    "qty_terjual": 7
  }
}
```

---

### 16. Get Report by Date Range
Retrieve sales report for a specific date range.

```bash
curl -X GET "http://localhost:6969/reports?start_date=2026-01-01&end_date=2026-12-31"
```

**Query Parameters:**
- `start_date`: Start date in YYYY-MM-DD format (optional)
- `end_date`: End date in YYYY-MM-DD format (optional)

**Response:**
```json
{
  "total_revenue": 17500,
  "total_transaksi": 6,
  "produk_terlaris": {
    "id": "8a046717-8407-4b22-b019-f7af47949c83",
    "nama": "Indomie Goreng",
    "qty_terjual": 7
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
  "date": "string (YYYY-MM-DD format)",
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
- Search functionality is available for both categories and products using the `search` query parameter
- Checkout response date field uses YYYY-MM-DD format (not ISO 8601)
