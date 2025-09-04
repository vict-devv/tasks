CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column ();

CREATE TRIGGER set_updated_at BEFORE
UPDATE ON categories FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column ();

CREATE TRIGGER set_updated_at BEFORE
UPDATE ON tasks FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column ();
