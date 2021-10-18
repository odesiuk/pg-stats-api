CREATE EXTENSION IF NOT EXISTS pg_stat_statements SCHEMA public CASCADE;

CREATE TABLE data
(
    id         serial      not null primary key,
    name       varchar(20) not null default 'unnamed',
    value      int         not null default 0,
    created_at timestamp   not null default now()
);

INSERT INTO data(name, value)
VALUES ('one', 20),
       ('one_1', 4214),
       ('gg', 140),
       ('two', 526);

UPDATE data
SET value=(SELECT value + 30 FROM pg_sleep(1))
WHERE value < 200;

DELETE
FROM data
WHERE value > (SELECT 100 FROM pg_sleep(1));

SELECT sum(value), pg_sleep(3)
FROM data;

SELECT *
FROM pg_stat_statements;