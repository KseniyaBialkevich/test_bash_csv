-- sudo -u postgres psql
-- \c csv_db
-- \d

CREATE DATABASE csv_db;

CREATE TABLE names (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(100) NOT NULL
    );

-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO kspsql;
-- GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO kspsql;