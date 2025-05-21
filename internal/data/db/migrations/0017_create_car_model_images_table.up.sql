-- migrations/0017_create_car_model_images_table.up.sql
CREATE TABLE car_model_images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE CASCADE,
    image_id UUID NOT NULL REFERENCES files(id) ON DELETE CASCADE,
    is_main_image BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_car_model_images_car_model_id ON car_model_images(car_model_id);
CREATE INDEX idx_car_model_images_image_id ON car_model_images(image_id);

-- Ensure only one main image per car model
CREATE OR REPLACE FUNCTION check_main_image()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.is_main_image THEN
        UPDATE car_model_images
        SET is_main_image = FALSE
        WHERE car_model_id = NEW.car_model_id AND id != NEW.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_check_main_image
AFTER INSERT OR UPDATE ON car_model_images
FOR EACH ROW EXECUTE FUNCTION check_main_image();

CREATE OR REPLACE FUNCTION update_car_model_images_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_images_modtime
BEFORE UPDATE ON car_model_images
FOR EACH ROW EXECUTE FUNCTION update_car_model_images_modtime();