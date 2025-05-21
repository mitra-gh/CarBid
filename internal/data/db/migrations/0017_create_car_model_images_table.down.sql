-- migrations/0017_create_car_model_images_table.down.sql
DROP TRIGGER IF EXISTS trigger_update_car_model_images_modtime ON car_model_images;
DROP TRIGGER IF EXISTS trigger_check_main_image ON car_model_images;
DROP FUNCTION IF EXISTS update_car_model_images_modtime;
DROP FUNCTION IF EXISTS check_main_image;
DROP INDEX IF EXISTS idx_car_model_images_car_model_id;
DROP INDEX IF EXISTS idx_car_model_images_image_id;
DROP TABLE IF EXISTS car_model_images;