-- migrations/0020_create_car_model_comments_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_model_comments_modtime ON car_model_comments;
DROP FUNCTION IF EXISTS update_car_model_comments_modtime;
DROP INDEX IF EXISTS idx_car_model_comments_car_model_id;
DROP INDEX IF EXISTS idx_car_model_comments_user_id;
DROP TABLE IF EXISTS car_model_comments;