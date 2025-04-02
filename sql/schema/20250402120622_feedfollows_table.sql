-- +goose Up
-- +goose StatementBegin
create table feed_follows
(
    id         uuid primary key default gen_random_uuid(),
    user_id    uuid      not null,
    feed_id    uuid      not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    constraint user_feed_unique
        unique (user_id, feed_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table feed_follows;
-- +goose StatementEnd
