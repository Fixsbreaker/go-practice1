-- 000002_create_categories_table.up.sql
CREATE TABLE categories (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT unique_user_category UNIQUE (user_id, name)
);
CREATE INDEX idx_categories_user_id ON categories(user_id);
