-- migrations/0012_create_files_table.up.sql
CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    directory VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_files_name ON files(name);

CREATE OR REPLACE FUNCTION update_files_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_files_modtime
BEFORE UPDATE ON files
FOR EACH ROW EXECUTE FUNCTION update_files_modtime();