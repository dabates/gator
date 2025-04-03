-- +goose Up
-- +goose StatementBegin
create table posts
(
    id           uuid primary key default gen_random_uuid(),
    title        text      not null,
    url          text      not null unique,
    description  text      not null,
    feed_id      uuid      not null,
    published_at timestamp not null,
    created_at   timestamp not null,
    updated_at   timestamp not null
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table posts;
-- +goose StatementEnd
