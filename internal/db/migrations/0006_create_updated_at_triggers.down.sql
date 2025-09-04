DROP TRIGGER IF EXISTS set_updated_at ON users;

DROP TRIGGER IF EXISTS set_updated_at ON categories;

DROP TRIGGER IF EXISTS set_updated_at ON tasks;

DROP FUNCTION IF EXISTS set_updated_at_column ();