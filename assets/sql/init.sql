CREATE SCHEMA IF NOT EXISTS anthive;

-- Set default search_path to schema
SET search_path TO anthive,public;

-- Creation of tables
CREATE TABLE IF NOT EXISTS antling (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50)
)