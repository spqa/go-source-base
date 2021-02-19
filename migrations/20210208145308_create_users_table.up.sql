create table users
(
    id         serial primary key,
    faculty_id bigint      not null references faculties (id),
    name       text,
    email      text unique not null,
    password   text,
    role       smallint,
    created_at timestamptz,
    updated_at timestamptz
);