DROP EXTENSION IF EXISTS "uuid-ossp";

alter table if exists users
    alter column if exists password type varchar(30),
    drop column if exists login;