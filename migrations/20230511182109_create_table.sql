q-- +goose Up
CREATE TABLE thread(
  id BIGINT(20) PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE post(
  id BIGINT(20) PRIMARY KEY,
  thread_id BIGINT(20) NOT NULL REFERENCES thread (id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  votes INT NOT NULL DEFAULT 0,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE comment(
  id BIGINT(20) PRIMARY KEY,
  post_id BIGINT(20) NOT NULL REFERENCES post (id) ON DELETE CASCADE,
  content Text NOT NULL,
  votes INT NOT NULL DEFAULT 0,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE thread;
DROP TABLE post;
DROP TABLE comment;
