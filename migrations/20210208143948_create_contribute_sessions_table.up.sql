create table contribute_sessions
(
    id                 serial primary key,
    open_time          timestamptz not null,
    closure_time       timestamptz not null,
    final_closure_time timestamptz not null,
    exported_assets    text,
    created_at         timestamptz,
    updated_at         timestamptz
);