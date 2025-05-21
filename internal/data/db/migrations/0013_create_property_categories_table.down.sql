-- migrations/0013_create_property_categories_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_property_categories_modtime ON property_categories;
DROP FUNCTION IF EXISTS update_property_categories_modtime;
DROP INDEX IF EXISTS idx_property_categories_name;
DROP TABLE IF EXISTS property_categories;