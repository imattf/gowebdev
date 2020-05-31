-- psql commands for prep

CREATE DATABASE bookstore;

CREATE USER bond WITH PASSWORD 'password';

GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;