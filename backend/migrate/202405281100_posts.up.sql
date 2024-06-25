USE training;

CREATE TABLE IF NOT EXISTS posts(
  id         INT          AUTO_INCREMENT PRIMARY KEY,
  user_id    INT          NOT NULL,
  title      VARCHAR(100) NOT NULL,
  body       TEXT         NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test2', '質問2\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test3', '質問3\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test4', '質問4\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test5', '質問5\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test6', '質問6\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test7', '質問7\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test8', '質問8\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test9', '質問9\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test10', '質問10\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test11', '質問11\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test12', '質問12\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test13', '質問13\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test14', '質問14\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test15', '質問15\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test16', '質問16\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test17', '質問17\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test18', '質問18\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test19', '質問19\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test20', '質問20\n改行');
