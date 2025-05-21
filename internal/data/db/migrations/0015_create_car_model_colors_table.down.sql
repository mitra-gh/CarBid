-- migrations/0015_create_car_model_colors_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_model_colors_modtime ON car_model_colors;
DROP FUNCTION IF EXISTS update_car_model_colors_modtime;
DROP INDEX IF EXISTS idx_car_model_colors_car_model_id;
DROP INDEX IF EXISTS idx_car_model_colors_color_id;
DROP TABLE IF EXISTS car_model_colors;