-- +goose Up
INSERT INTO users (login, password_hash)
VALUES ('admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918');

-- +goose Down
DELETE FROM users WHERE login = 'admin';
