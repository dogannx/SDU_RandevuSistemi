CREATE TABLE IF NOT EXISTS teachers (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    subjects    TEXT[] NOT NULL,
    bio         TEXT
);
