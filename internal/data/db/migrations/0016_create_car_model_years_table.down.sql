-- migrations/0016_create_car_model_years_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_model_years_modtime ON car_model_years;
DROP FUNCTION IF EXISTS update_car_model_years_modtime;
DROP INDEX IF EXISTS idx_car_model_years_car_model_id;
DROP INDEX IF EXISTS idx_car_model_years_persian_year_id;
DROP TABLE IF EXISTS car_model_years;