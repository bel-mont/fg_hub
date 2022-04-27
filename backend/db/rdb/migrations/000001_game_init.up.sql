CREATE TABLE IF NOT EXISTS games(
   id uuid PRIMARY KEY,
   created_at timestamp,
   updated_at timestamp,
   name VARCHAR (300)
);