CREATE TABLE IF NOT EXISTS users (
	id       SERIAL NOT NULL,
	nik   varchar(255) NOT NULL UNIQUE,
	password varchar(255) NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);