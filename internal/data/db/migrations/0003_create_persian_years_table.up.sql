-- migrations/0003_create_persian_years_table.up.sql
CREATE TABLE persian_years (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    persian_title VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,
    start_at TIMESTAMPTZ  NOT NULL,
    end_at TIMESTAMPTZ  NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

ALTER TABLE persian_years ADD CONSTRAINT uq_persian_years_year UNIQUE (year);

CREATE INDEX idx_persian_years_year ON persian_years(year);

CREATE OR REPLACE FUNCTION update_persian_years_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_persian_years_modtime
BEFORE UPDATE ON persian_years
FOR EACH ROW EXECUTE FUNCTION update_persian_years_modtime();