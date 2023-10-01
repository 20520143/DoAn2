-- +migrate Up
CREATE TABLE users (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL 
);
-- +migrate Down
DROP TABLE users;
