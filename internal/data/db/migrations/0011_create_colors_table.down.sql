-- migrations/0011_create_colors_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_colors_modtime ON colors;
DROP FUNCTION IF EXISTS update_colors_modtime;
DROP INDEX IF EXISTS idx_colors_name;
ALTER TABLE colors DROP CONSTRAINT IF EXISTS uq_colors_hex_code;
DROP TABLE IF EXISTS colors;