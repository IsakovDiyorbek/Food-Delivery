CREATE TYPE statusenum AS ENUM ('Pending', 'Assigned', 'Picked Up', 'Out for Delivery', 'Delivered', 'Cancelled');

CREATE TABLE products (
    id uuid not null PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE orders (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    courier_id uuid NOT NULL,
    status statusenum,
    total_amount DECIMAL(10, 2) NOT NULL,
    delivery_address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0 
);

CREATE TABLE cart (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    product_id uuid REFERENCES products(id),
    quantity INTEGER NOT NULL,
    options JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Order Items Table
CREATE TABLE order_items (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid REFERENCES orders(id),
    product_id uuid REFERENCES products(id),
    quantity INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);



-- Notifications Table
CREATE TABLE notifications (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid not null,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tasks Table
CREATE TABLE tasks (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL,
    user_id_assigned_to INTEGER not null,
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

-- Courier Locations Table
CREATE TABLE courier_locations (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    courier_id uuid not null,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    start_time TIMESTAMP not null,
    end_time TIMESTAMP not null,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0 
);


