CREATE TABLE IF NOT EXISTS product (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR NOT NULL,
    stock int,
    created_at TIMESTAMP DEFAULT NOW()
);