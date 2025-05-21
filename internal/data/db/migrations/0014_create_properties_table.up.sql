-- migrations/0014_create_properties_table.up.sql
CREATE TABLE properties (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    category_id UUID NOT NULL REFERENCES property_categories(id),
    description TEXT NOT NULL,
    data_type VARCHAR(50) NOT NULL,
    unit VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_properties_name ON properties(name);
CREATE INDEX idx_properties_category_id ON properties(category_id);

CREATE OR REPLACE FUNCTION update_properties_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_properties_modtime
BEFORE UPDATE ON properties
FOR EACH ROW EXECUTE FUNCTION update_properties_modtime();