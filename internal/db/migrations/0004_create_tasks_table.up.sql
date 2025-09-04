CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  status SMALLINT NOT NULL DEFAULT 0, -- 0: pending, 1: in-progress, 2: completed, 3: cancelled
  priority SMALLINT NOT NULL DEFAULT 0, -- 0: low, 1: medium, 2: high, 3: urgent
  due_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE
);
