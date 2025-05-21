-- migrations/0014_create_properties_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_properties_modtime ON properties;
DROP FUNCTION IF EXISTS update_properties_modtime;
DROP INDEX IF EXISTS idx_properties_name;
DROP INDEX IF EXISTS idx_properties_category_id;
DROP TABLE IF EXISTS properties;