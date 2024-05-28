USE training;

CREATE TABLE IF NOT EXISTS hello_worlds(
  lang    VARCHAR(2) NOT NULL PRIMARY KEY,
  message VARCHAR(40) NOT NULL
);
INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');
