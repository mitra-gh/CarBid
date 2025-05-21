-- migrations/0001_create_countries_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_countries_modtime ON countries;
DROP FUNCTION IF EXISTS update_countries_modtime;
DROP INDEX IF EXISTS idx_countries_name;
DROP TABLE IF EXISTS countries;