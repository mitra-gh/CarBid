-- migrations/0008_create_car_types_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_types_modtime ON car_types;
DROP FUNCTION IF EXISTS update_car_types_modtime;
DROP INDEX IF EXISTS idx_car_types_name;
DROP TABLE IF EXISTS car_types;