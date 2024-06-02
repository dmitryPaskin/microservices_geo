DROP TABLE IF EXISTS users;

CREATE TABLE users(
                      id       SERIAL PRIMARY KEY,
                      login    VARCHAR(255),
                      password VARCHAR(255)
);
