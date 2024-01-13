create type user_role_enum as enum ('admin', 'customer');

create table drivers (
    id uuid primary key,
    full_name text,
    phone text not null
);

create table cars (
    id uuid primary key,
    year int
    model text not null,
    brand text not null,
    driver_id uuid references drivers(id)
);

create table users (
    id uuid primary key not null,
    full_name varchar(30),
    phone varchar(30) unique not null,
    password varchar(30) not null,
    user_role user_role_enum not null,
    cash int
);