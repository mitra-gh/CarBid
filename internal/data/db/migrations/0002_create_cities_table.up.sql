-- migrations/0002_create_cities_table.up.sql
CREATE TABLE cities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    country_id UUID NOT NULL REFERENCES countries(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_cities_country_id ON cities(country_id);
CREATE INDEX idx_cities_name ON cities(name);

CREATE OR REPLACE FUNCTION update_cities_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_cities_modtime
BEFORE UPDATE ON cities
FOR EACH ROW EXECUTE FUNCTION update_cities_modtime();