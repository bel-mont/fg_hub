
ALTER TABLE games
    ALTER COLUMN id SET DEFAULT NULL,
    ALTER COLUMN created_at SET DEFAULT NULL,
    ALTER COLUMN updated_at SET DEFAULT NULL;
    -- ADD COLUMN slug VARCHAR (50) NOT NULL; -- DO NOT DROP THE SLUG NO MATTER WHAT! It is a useful key!


ALTER TABLE characters
    ALTER COLUMN id SET DEFAULT NULL,
    ALTER COLUMN created_at SET DEFAULT NULL,
    ALTER COLUMN updated_at SET DEFAULT NULL;

