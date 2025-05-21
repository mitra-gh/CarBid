-- migrations/0015_create_car_model_colors_table.up.sql
CREATE TABLE car_model_colors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE CASCADE,
    color_id UUID NOT NULL REFERENCES colors(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(car_model_id, color_id)
);

CREATE INDEX idx_car_model_colors_car_model_id ON car_model_colors(car_model_id);
CREATE INDEX idx_car_model_colors_color_id ON car_model_colors(color_id);

CREATE OR REPLACE FUNCTION update_car_model_colors_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_colors_modtime
BEFORE UPDATE ON car_model_colors
FOR EACH ROW EXECUTE FUNCTION update_car_model_colors_modtime();