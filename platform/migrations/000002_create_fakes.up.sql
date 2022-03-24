BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Asia/Jakarta";

-- Create users table
CREATE TABLE fakes (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    type VARCHAR (255) NOT NULL,
    code VARCHAR (255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    nakes_count NUMERIC CHECK (nakes_count > 0),
    created_by UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP 
);

-- Add indexes
CREATE INDEX name_fakes ON fakes(name);

CREATE INDEX id_fakes ON fakes(id);

CREATE INDEX code_fakes ON fakes(code);

COMMIT