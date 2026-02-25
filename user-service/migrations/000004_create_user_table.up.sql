CREATE TABLE app_user (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(500) NOT NULL,

    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    phone VARCHAR(20),

    user_status_id INT NOT NULL,
    user_role_id INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL
);

CREATE UNIQUE INDEX uq_app_user_email 
ON app_user(email);

CREATE INDEX idx_app_user_status_id 
ON app_user(user_status_id);