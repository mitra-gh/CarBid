-- migrations/0018_create_car_model_price_histories_table.up.sql
CREATE TABLE car_model_price_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_year_id UUID NOT NULL REFERENCES car_model_years(id) ON DELETE CASCADE,
    price FLOAT NOT NULL DEFAULT 0,
    price_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_car_model_price_histories_car_model_year_id ON car_model_price_histories(car_model_year_id);
CREATE INDEX idx_car_model_price_histories_price_at ON car_model_price_histories(price_at);

CREATE OR REPLACE FUNCTION update_car_model_price_histories_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_price_histories_modtime
BEFORE UPDATE ON car_model_price_histories
FOR EACH ROW EXECUTE FUNCTION update_car_model_price_histories_modtime();