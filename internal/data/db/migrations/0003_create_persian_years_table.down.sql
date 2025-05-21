-- migrations/0003_create_persian_years_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_persian_years_modtime ON persian_years;
DROP FUNCTION IF EXISTS update_persian_years_modtime;
DROP INDEX IF EXISTS idx_persian_years_year;
ALTER TABLE persian_years DROP CONSTRAINT IF EXISTS uq_persian_years_year;
DROP TABLE IF EXISTS persian_years;