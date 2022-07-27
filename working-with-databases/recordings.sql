CREATE DATABASE recordings

CREATE TABLE album(
id INT IDENTITY(1,1) PRIMARY KEY,
title NVARCHAR(255),
artist NVARCHAR(255),
price INT
)
GO
INSERT INTO album (title,artist,price) VALUES ('Hari Yang Cerah','Peterpan',50000)
INSERT INTO album (title,artist,price) VALUES ('Sebuah Nama Sebuah Cerita','Peterpan',50000)
INSERT INTO album (title,artist,price) VALUES ('Bintang Di surga','Peterpan',50000)
GO
