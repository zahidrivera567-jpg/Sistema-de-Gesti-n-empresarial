CREATE DATABASE IF NOT EXISTS gestion_libros;
USE gestion_libros;
-- TABLA: usuarios
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    correo VARCHAR(100) NOT NULL UNIQUE,
    contrasena VARCHAR(255) NOT NULL,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- TABLA: categorias
CREATE TABLE categorias (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE
);
-- TABLA: libros
CREATE TABLE libros (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(200) NOT NULL,
    autor VARCHAR(150) NOT NULL,
    anio YEAR,
    descripcion TEXT,
    categoria_id INT,
    usuario_id INT,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE SET NULL ON UPDATE CASCADE
);
-- INSERTAR CATEGORÍAS BÁSICAS
INSERT INTO categorias (nombre) VALUES
('Ciencia'),
('Historia'),
('Literatura'),
('Tecnología'),
('Arte');



-- EJEMPLOS !!NO EJECUTAR!!
-- INSERTAR USUARIO DE PRUEBA
INSERT INTO usuarios (nombre, correo, contrasena)
VALUES ('Administrador', 'admin@biblioteca.com', 'admin123');
-- INSERTAR LIBRO DE PRUEBA
INSERT INTO libros (titulo, autor, anio, descripcion, categoria_id, usuario_id)
VALUES ('El origen de las especies', 'Charles Darwin', 1859, 'Obra fundamental sobre la teoría de la evolución.', 1, 1);
