create table article_versions
(
    id            serial primary key,
    hash          text        not null,
    article_id    bigint references articles (id),
    link_original text        not null,
    link_pdf      text,
    created_at    timestamptz not null
)