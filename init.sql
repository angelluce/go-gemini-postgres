CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT
);

INSERT INTO items (name, description) VALUES
                                          ('Laptop Gaming', 'Una potente laptop diseñada para juegos con gráficos RTX 4080 y 32GB de RAM DDR5.'),
                                          ('Teclado Mecánico RGB', 'Teclado con switches Cherry MX Brown, retroiluminación RGB personalizable y diseño ergonómico.'),
                                          ('Mouse Inalámbrico Ergonómico', 'Mouse de alta precisión con sensor óptico de 16000 DPI y conectividad inalámbrica 2.4 GHz y Bluetooth.');