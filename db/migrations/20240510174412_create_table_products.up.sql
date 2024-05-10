CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(30) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    category VARCHAR(255) NOT NULL CHECK (category IN ('Clothing', 'Accessories', 'Footwear', 'Beverages')),
    image_url VARCHAR(255) NOT NULL,
    notes VARCHAR(200) NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 1),
    stock INTEGER NOT NULL CHECK (stock >= 0 AND stock <= 100000),
    location VARCHAR(200) NOT NULL,
    is_available BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)