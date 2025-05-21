-- migrations/0010_create_car_models_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_models_modtime ON car_models;
DROP FUNCTION IF EXISTS update_car_models_modtime;
DROP INDEX IF EXISTS idx_car_models_name;
DROP INDEX IF EXISTS idx_car_models_company_id;
DROP INDEX IF EXISTS idx_car_models_car_type_id;
DROP INDEX IF EXISTS idx_car_models_gearbox_id;
DROP TABLE IF EXISTS car_models;