CREATE TABLE expenses (
                          id SERIAL PRIMARY KEY,
                          user_id     INTEGER NOT NULL REFERENCES users(id),
                          category_id INTEGER NOT NULL REFERENCES categories(id),
                          amount   NUMERIC(12,2) NOT NULL CHECK (amount > 0),
                          currency CHAR(3) NOT NULL,
                          spent_at TIMESTAMP NOT NULL,
                          created_at TIMESTAMP DEFAULT NOW(),
                          note TEXT
);
CREATE INDEX idx_expenses_user_id    ON expenses(user_id);
CREATE INDEX idx_expenses_user_spent ON expenses(user_id, spent_at);
