BEGIN;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Asia/Jakarta";

-- Create users table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    email VARCHAR (255) NOT NULL UNIQUE,
    username VARCHAR (255) NOT NULL UNIQUE,
    password VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Add indexes
CREATE INDEX email_users ON users(email);

CREATE INDEX id_users ON users(id);

CREATE INDEX username_users ON users(username);

COMMIT