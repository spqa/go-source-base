create table articles
(
    id          serial primary key,
    title       text not null,
    description text,
    created_at  timestamptz,
    updated_at  timestamptz
);