-- migrations/0002_create_cities_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_cities_modtime ON cities;
DROP FUNCTION IF EXISTS update_cities_modtime;
DROP INDEX IF EXISTS idx_cities_country_id;
DROP INDEX IF EXISTS idx_cities_name;
DROP TABLE IF EXISTS cities;