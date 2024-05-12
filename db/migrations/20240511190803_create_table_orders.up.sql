CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    customer_id VARCHAR(255) NOT NULL,
    paid INT NOT NULL CHECK (paid >= 1),
    change INT NOT NULL CHECK (change >= 0),
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_details (
    order_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 1),
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    created_at TIMESTAMP NOT NULL
);