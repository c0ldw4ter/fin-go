-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    balance DECIMAL(10, 2) DEFAULT 0
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_user_id INT,
    to_user_id INT,
    amount DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (from_user_id) REFERENCES users(id),
    FOREIGN KEY (to_user_id) REFERENCES users(id)
);
-- +goose Down
DROP TABLE transactions;
DROP TABLE users;