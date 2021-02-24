create table users
(
    id         serial primary key,
    faculty_id bigint      default null references faculties (id),
    name       text,
    email      text unique not null,
    password   text,
    role       text,
    created_at timestamptz,
    updated_at timestamptz
);