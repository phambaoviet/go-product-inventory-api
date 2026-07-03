## Environment Variables

Create a `.env` file based on `.env.example`.

```env
API_KEY=your-api-key

DB_HOST=localhost
DB_PORT=5433
DB_NAME=product_inventory_db
DB_USER=root
DB_PASSWORD=your-password
DB_PASSWORD_ENCODED=your-url-encoded-password
DB_SSLMODE=disable
```

Note: If your database password contains special characters, use the URL-encoded version in `DB_PASSWORD_ENCODED` for migration commands.

Example:

```text
viet@123456 -> viet%40123456
```

## Getting Started

### 1. Clone the repository

```bash
git clone <your-repository-url>
cd go-product-inventory-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create environment file

```bash
cp .env.example .env
```

Then update `.env` with your local values.

### 4. Start PostgreSQL with Docker Compose

```bash
docker compose up -d
```

### 5. Run database migrations

```bash
make migrate-up
```

### 6. Start the server

```bash
go run .
```

The server will run at:

```text
http://localhost:8080
```

## API Authentication

All product routes require an API key header:

```text
X-API-Key: your-api-key
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/products` | Get all products |
| GET | `/api/products/:id` | Get product by ID |
| POST | `/api/products` | Create a new product |
| PUT | `/api/products/:id` | Update product by ID |
| DELETE | `/api/products/:id` | Delete product by ID |

## Query Filters

Filter by product name:

```text
GET /api/products?name=keyboard
```

Filter by product slug:

```text
GET /api/products?slug=mechanical-keyboard
```

## Example Requests

### Get all products

```bash
curl http://localhost:8080/api/products \
-H "X-API-Key: your-api-key"
```

### Create product

```bash
curl -X POST http://localhost:8080/api/products \
-H "Content-Type: application/json" \
-H "X-API-Key: your-api-key" \
-d '{
  "name": "Gaming Mouse",
  "price": 29.99,
  "quantity": 10
}'
```

### Update product

```bash
curl -X PUT http://localhost:8080/api/products/1 \
-H "Content-Type: application/json" \
-H "X-API-Key: your-api-key" \
-d '{
  "name": "Gaming Mouse Updated",
  "price": 35.99,
  "quantity": 15
}'
```

### Delete product

```bash
curl -X DELETE http://localhost:8080/api/products/1 \
-H "X-API-Key: your-api-key"
```

## Migration Commands

Create a migration:

```bash
make migrate-create name=create_products_table
```

Run migrations:

```bash
make migrate-up
```

Rollback one migration:

```bash
make migrate-down
```

Force migration version:

```bash
make migrate-force version=1
```

## Current Status

This project currently supports product CRUD operations with PostgreSQL and API key authentication.