-- migrations/0019_create_car_model_properties_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_model_properties_modtime ON car_model_properties;
DROP FUNCTION IF EXISTS update_car_model_properties_modtime;
DROP INDEX IF EXISTS idx_car_model_properties_car_model_id;
DROP INDEX IF EXISTS idx_car_model_properties_property_id;
DROP TABLE IF EXISTS car_model_properties;