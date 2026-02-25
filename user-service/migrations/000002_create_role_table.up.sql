CREATE TABLE user_role (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255)
);

INSERT INTO user_role (code, description) VALUES
('user', 'Normal user'),
('admin', 'Administrator'),
('moderator', 'Moderator');