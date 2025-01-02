DROP SCHEMA IF EXISTS SAMPLE_DB;
CREATE SCHEMA SAMPLE_DB;
USE SAMPLE_DB;

DROP TABLE IF EXISTS user;

CREATE TABLE user (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);

INSERT INTO user (name, email) VALUES ('John Doe', 'John@example.com');
INSERT INTO user (name, email) VALUES ('Kathy Smith', 'Kathy@example.com');