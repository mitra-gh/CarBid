-- migrations/0009_create_companies_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_companies_modtime ON companies;
DROP FUNCTION IF EXISTS update_companies_modtime;
DROP INDEX IF EXISTS idx_companies_name;
DROP INDEX IF EXISTS idx_companies_country_id;
DROP TABLE IF EXISTS companies;