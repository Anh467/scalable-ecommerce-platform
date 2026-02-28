CREATE TABLE user_status (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255)
);

INSERT INTO user_status (code, description) VALUES
('active', 'User is active'),
('inactive', 'User is inactive'),
('banned', 'User is banned');