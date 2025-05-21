-- migrations/0007_create_gearboxes_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_gearboxes_modtime ON gearboxes;
DROP FUNCTION IF EXISTS update_gearboxes_modtime;
DROP INDEX IF EXISTS idx_gearboxes_name;
DROP TABLE IF EXISTS gearboxes;