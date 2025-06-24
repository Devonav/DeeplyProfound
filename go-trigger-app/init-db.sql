CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE technologies (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  name    VARCHAR(255),
  details VARCHAR(1024)
);