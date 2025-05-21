-- migrations/0007_create_gearboxes_table.up.sql
CREATE TABLE gearboxes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_gearboxes_name ON gearboxes(name);

CREATE OR REPLACE FUNCTION update_gearboxes_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_gearboxes_modtime
BEFORE UPDATE ON gearboxes
FOR EACH ROW EXECUTE FUNCTION update_gearboxes_modtime();