-- migrations/0006_create_user_roles_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_user_roles_modtime ON user_roles;
DROP FUNCTION IF EXISTS update_user_roles_modtime;
DROP INDEX IF EXISTS idx_user_roles_user_id;
DROP INDEX IF EXISTS idx_user_roles_role_id;
DROP TABLE IF EXISTS user_roles;