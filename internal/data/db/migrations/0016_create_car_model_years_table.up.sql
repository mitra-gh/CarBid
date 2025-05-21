-- migrations/0016_create_car_model_years_table.up.sql
CREATE TABLE car_model_years (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE CASCADE,
    persian_year_id UUID NOT NULL REFERENCES persian_years(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(car_model_id, persian_year_id)
);

CREATE INDEX idx_car_model_years_car_model_id ON car_model_years(car_model_id);
CREATE INDEX idx_car_model_years_persian_year_id ON car_model_years(persian_year_id);

CREATE OR REPLACE FUNCTION update_car_model_years_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_years_modtime
BEFORE UPDATE ON car_model_years
FOR EACH ROW EXECUTE FUNCTION update_car_model_years_modtime();