-- migrations/0008_create_car_types_table.up.sql
CREATE TABLE car_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_car_types_name ON car_types(name);

CREATE OR REPLACE FUNCTION update_car_types_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_types_modtime
BEFORE UPDATE ON car_types
FOR EACH ROW EXECUTE FUNCTION update_car_types_modtime();