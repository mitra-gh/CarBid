-- migrations/0005_create_users_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_users_modtime ON users;
DROP FUNCTION IF EXISTS update_users_modtime;
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_mobile_number;
DROP TABLE IF EXISTS users;