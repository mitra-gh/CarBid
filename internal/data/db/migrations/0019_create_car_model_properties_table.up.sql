-- migrations/0019_create_car_model_properties_table.up.sql
CREATE TABLE car_model_properties (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE CASCADE,
    property_id UUID NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(car_model_id, property_id)
);

CREATE INDEX idx_car_model_properties_car_model_id ON car_model_properties(car_model_id);
CREATE INDEX idx_car_model_properties_property_id ON car_model_properties(property_id);

CREATE OR REPLACE FUNCTION update_car_model_properties_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_properties_modtime
BEFORE UPDATE ON car_model_properties
FOR EACH ROW EXECUTE FUNCTION update_car_model_properties_modtime();