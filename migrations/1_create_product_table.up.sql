CREATE TABLE IF NOT EXISTS product (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    stock int,
    created_at TIMESTAMP DEFAULT NOW()
);