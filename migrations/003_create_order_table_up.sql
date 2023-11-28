CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user jsonb NOT NULL,
    status VARCHAR(255) NOT NULL,
    order_items jsonb NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);
