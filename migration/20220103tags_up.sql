CREATE TABLE IF NOT EXISTS tags (
	tag   varchar(255) NOT NULL UNIQUE,
	description varchar(255) NOT NULL,
	user_id int NOT NULL,
	CONSTRAINT tags_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);