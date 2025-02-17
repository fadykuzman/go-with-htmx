
CREATE TABLE IF NOT EXISTS dogs (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    breed VARCHAR(100) NOT NULL
);

INSERT INTO dogs (name, breed) VALUES
    ('Max', 'Labrador'),
    ('Luna', 'German Shepherd'),
    ('Bella', 'Golden Retriever');
