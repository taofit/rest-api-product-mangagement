CREATE TABLE IF NOT EXISTS product_price (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_id uuid NOT NULL,
    price int NOT NULL,
    currency VARCHAR NOT NULL,
    UNIQUE(product_id,currency),
    FOREIGN KEY(product_id) REFERENCES product(id) ON DELETE CASCADE
);