USE training;

CREATE TABLE IF NOT EXISTS comments(
  id         INT          AUTO_INCREMENT PRIMARY KEY,
  user_id    INT          NOT NULL,
  post_id    INT          NOT NULL,
  body       TEXT         NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (post_id) REFERENCES posts (id)
);
INSERT INTO comments (user_id, post_id, body) VALUES (1, 1, 'コメント一号');
INSERT INTO comments (user_id, post_id, body) VALUES (2, 2, 'コメント二号');
