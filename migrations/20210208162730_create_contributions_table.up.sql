create table contributions
(
    id                    serial primary key,
    user_id               bigint  not null references users (id),
    contribute_session_id bigint  not null references contribute_sessions (id),
    article_id            bigint  not null references articles (id),
    images                jsonb,
    selected              boolean not null default false,
    created_at            timestamptz,
    updated_at            timestamptz
)