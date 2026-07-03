CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    price DOUBLE PRECISION NOT NULL CHECK (price > 0),
quantity INT NOT NULL CHECK (quantity >= 0)
);