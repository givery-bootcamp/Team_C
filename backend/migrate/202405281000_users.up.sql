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

INSERT INTO users (name, password) VALUES ('taro', '$2a$10$ua2lNLVfRx0XlImwQxZx7eOe./3iFQuwmfYx1ylkwN5887w341IM6');
INSERT INTO users (name, password) VALUES ('hanako', '$2a$10$91jrQap/pHIkKXGwp4UtT.EMw5ac3ZEC53M1EECNNRde2/t3QLt4y');
