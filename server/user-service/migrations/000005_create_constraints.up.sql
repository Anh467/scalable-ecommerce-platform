ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_status
FOREIGN KEY (user_status_id)
REFERENCES user_status(id);

ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_role
FOREIGN KEY (user_role_id)
REFERENCES user_role(id);