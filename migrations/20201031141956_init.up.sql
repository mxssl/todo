CREATE TABLE items (
	id serial PRIMARY KEY,
	description TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE
);
