
CREATE DATABASE IF NOT EXISTS article;
USE article;

CREATE TABLE IF NOT EXISTS posts (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(200),
    Content TEXT,
    Category VARCHAR(100),
    Created_date TIMESTAMP,
    Updated_date TIMESTAMP,
    Status VARCHAR(100) CHECK (Status IN ('Publish', 'Draft', 'Thrash'))
);