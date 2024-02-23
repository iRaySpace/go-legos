CREATE TABLE employees (
    id          uuid    DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name  VARCHAR NOT NULL,
    last_name   VARCHAR NOT NULL
);