CREATE TABLE IF NOT EXISTS product_price (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    product_id uuid NOT NULL,
    price NUMERIC NOT NULL,
    currency VARCHAR NOT NULL,
    UNIQUE(product_id,currency),
    FOREIGN KEY(product_id) REFERENCES product(id) ON DELETE CASCADE
);