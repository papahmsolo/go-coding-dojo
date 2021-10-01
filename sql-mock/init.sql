
CREATE TABLE IF NOT EXISTS file (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS request (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    file_id uuid REFERENCES file(id)
);
