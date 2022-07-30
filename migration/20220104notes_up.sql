CREATE TABLE IF NOT EXISTS notes (
	id          varchar(255) NOT NULL UNIQUE,
	category 	categories NOT NULL,
	user_id     int NOT NULL,
	description text NOT NULL,
	links       varchar(255)[],
	tags        varchar(255)[],
	sources     text,
	is_first	boolean NOT NULL default FALSE,
	created_at  TIMESTAMPTZ NOT NULL,
	last_updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT notes_pkey PRIMARY KEY (id),
	CONSTRAINT notes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

comment on column notes.id is 'Уникальный номер карточки.';
comment on column notes.category is 'Категория (тип) карточки берется из типа categories.';
comment on column notes.user_id is 'ID пользователя составевшего карточку.';
comment on column notes.description is 'Наполнение карточки - мысли или занания записанные своими словами, для первых карточек в векторе о чем будет этот вектор.';
comment on column notes.links is 'Связи с другими карточками. Основа системы.';
comment on column notes.tags is 'Теги для объедения карточек по тематикам.';
comment on column notes.sources is 'Сылки на источники. Особенно важно для карточек из раздела знаний.';
comment on column notes.is_first is 'Флаг является ли карточка первой в векторе.';
comment on column notes.created_at is 'Время создания карточки.';
comment on column notes.last_updated_at is 'Время последнего обновления карточки.';
