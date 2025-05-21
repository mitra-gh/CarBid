-- migrations/0021_create_current_car_prices_view.up.sql
CREATE OR REPLACE VIEW current_car_prices AS
SELECT 
    cm.id AS car_model_id,
    cm.name AS car_model_name,
    c.name AS company_name,
    py.year AS persian_year,
    cmph.price AS current_price,
    cmph.price_at AS price_updated_at
FROM car_models cm
JOIN companies c ON cm.company_id = c.id
JOIN car_model_years cmy ON cm.id = cmy.car_model_id
JOIN persian_years py ON cmy.persian_year_id = py.id
JOIN (
    SELECT 
        car_model_year_id,
        price,
        price_at,
        ROW_NUMBER() OVER (PARTITION BY car_model_year_id ORDER BY price_at DESC) AS rn
    FROM car_model_price_histories
) cmph ON cmy.id = cmph.car_model_year_id AND cmph.rn = 1;