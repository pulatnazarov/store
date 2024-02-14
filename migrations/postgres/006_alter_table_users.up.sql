CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

alter table if exists users
    alter column password type text,
    add column if not exists login varchar(40) unique not null default uuid_generate_v4();
