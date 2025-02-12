\c postgres;

CREATE DATABASE dogs;

\c dogs;

CREATE TABLE dogs (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(50) NOT NULL,
	breed VARCHAR(50) NOT NULL
);


INSERT INTO dogs (name, breed) VALUES
	('Max', 'Golden Retriever'),
    ('Luna', 'German Shepherd');
