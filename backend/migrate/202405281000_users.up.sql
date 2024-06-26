USE training;

CREATE TABLE IF NOT EXISTS users(
  id         INT          AUTO_INCREMENT PRIMARY KEY,
  name       VARCHAR(40)  NOT NULL,
  password   VARCHAR(100) NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL,
  UNIQUE KEY(name)
);

INSERT INTO users (name, password) VALUES ('taro', 'password');
INSERT INTO users (name, password) VALUES ('hanako', 'PASSWORD');
