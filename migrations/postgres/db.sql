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
