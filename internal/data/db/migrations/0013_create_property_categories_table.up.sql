-- migrations/0013_create_property_categories_table.up.sql
CREATE TABLE property_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_property_categories_name ON property_categories(name);

CREATE OR REPLACE FUNCTION update_property_categories_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_property_categories_modtime
BEFORE UPDATE ON property_categories
FOR EACH ROW EXECUTE FUNCTION update_property_categories_modtime();