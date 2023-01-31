CREATE TABLE IF NOT EXISTS main.accounts
(
    id     SERIAL PRIMARY KEY,
    wallet TEXT,
    amount FLOAT
);