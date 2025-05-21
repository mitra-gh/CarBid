-- migrations/0010_create_car_models_table.up.sql
CREATE TABLE car_models (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    company_id UUID NOT NULL REFERENCES companies(id),
    car_type_id UUID NOT NULL REFERENCES car_types(id),
    gearbox_id UUID NOT NULL REFERENCES gearboxes(id),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_car_models_name ON car_models(name);
CREATE INDEX idx_car_models_company_id ON car_models(company_id);
CREATE INDEX idx_car_models_car_type_id ON car_models(car_type_id);
CREATE INDEX idx_car_models_gearbox_id ON car_models(gearbox_id);

CREATE OR REPLACE FUNCTION update_car_models_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_models_modtime
BEFORE UPDATE ON car_models
FOR EACH ROW EXECUTE FUNCTION update_car_models_modtime();