-- migrations/0012_create_files_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_files_modtime ON files;
DROP FUNCTION IF EXISTS update_files_modtime;
DROP INDEX IF EXISTS idx_files_name;
DROP TABLE IF EXISTS files;