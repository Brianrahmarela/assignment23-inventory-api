create database inventaris;

use inventaris;

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    category VARCHAR(50)
);

CREATE TABLE inventory (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    quantity INT,
    location VARCHAR(100),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    quantity INT,
    order_date DATE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO products (name, description, price, category) VALUES
('Laptop Dell', 'Laptop gaming', 20000000, 'Elektronik'),
('Bangku Kantor pexio', 'Bangku karyawan', 1000000, 'Peralatan'),
('PC Mac', 'Komputer kantor', 30000000, 'Elektronik');

INSERT INTO inventory (product_id, quantity, location) VALUES
(1, 1, 'Gudang Jakarta'),
(2, 10, 'Gudang Bekasi'),
(3, 1, 'Gudang Jakarta');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 2, '2025-06-16'),
(2, 3, '2025-06-10'),
(2, 7, '2025-06-10'),
(3, 5, '2025-06-11');

select * from orders;

-- Cek Semua produk dan stoknya
SELECT p.name, i.quantity, i.location
FROM products p
JOIN inventory i ON p.id = i.product_id;

-- Total pesanan untuk setiap produk
SELECT p.name, SUM(o.quantity) as total_ordered
FROM orders o
JOIN products p ON o.product_id = p.id
GROUP BY p.name;

-- Stok di lokasi tertentu
SELECT p.name, i.quantity
FROM inventory i
JOIN products p ON i.product_id = p.id
WHERE i.location = 'Gudang Bekasi';