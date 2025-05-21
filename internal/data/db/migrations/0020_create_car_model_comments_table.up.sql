-- migrations/0020_create_car_model_comments_table.up.sql
CREATE TABLE car_model_comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_car_model_comments_car_model_id ON car_model_comments(car_model_id);
CREATE INDEX idx_car_model_comments_user_id ON car_model_comments(user_id);

CREATE OR REPLACE FUNCTION update_car_model_comments_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_car_model_comments_modtime
BEFORE UPDATE ON car_model_comments
FOR EACH ROW EXECUTE FUNCTION update_car_model_comments_modtime();