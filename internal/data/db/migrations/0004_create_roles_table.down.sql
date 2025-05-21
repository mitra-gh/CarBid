-- migrations/0004_create_roles_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_roles_modtime ON roles;
DROP FUNCTION IF EXISTS update_roles_modtime;
DROP INDEX IF EXISTS idx_roles_name;
DROP TABLE IF EXISTS roles;