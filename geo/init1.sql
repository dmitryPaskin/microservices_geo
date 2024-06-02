DROP TABLE IF EXISTS address_data;
DROP TABLE IF EXISTS geo_data;

CREATE TABLE address_data(
                             id      SERIAL PRIMARY KEY,
                             address VARCHAR(255),
                             data    TEXT
);

CREATE TABLE geo_data(
                         id SERIAL PRIMARY KEY,
                         geo VARCHAR(255),
                         data TEXT
);