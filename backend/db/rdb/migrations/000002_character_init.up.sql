-- Sample only, redo/delete later
CREATE TABLE IF NOT EXISTS characters(
   id uuid PRIMARY KEY,
   created_at timestamp,
   updated_at timestamp,
   name VARCHAR (300),
   game_id uuid references games (id)
);