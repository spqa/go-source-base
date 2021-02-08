create extension if not exists "uuid-ossp";
create table comments
(
    id              uuid primary key default uuid_generate_v4(),
    user_id         bigint  not null references users (id),
    contribution_id bigint  not null references contributions (id),
    content         text    not null,
    resolved        boolean not null default false,
    created_at      timestamptz,
    updated_at      timestamptz
)