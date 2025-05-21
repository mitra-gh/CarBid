-- migrations/0011_create_colors_table.up.sql
CREATE TABLE colors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    hex_code VARCHAR(7) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

ALTER TABLE colors ADD CONSTRAINT uq_colors_hex_code UNIQUE (hex_code);

CREATE INDEX idx_colors_name ON colors(name);

CREATE OR REPLACE FUNCTION update_colors_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_colors_modtime
BEFORE UPDATE ON colors
FOR EACH ROW EXECUTE FUNCTION update_colors_modtime();