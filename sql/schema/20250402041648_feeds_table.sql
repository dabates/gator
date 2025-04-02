-- +goose Up
-- +goose StatementBegin
create table feeds
(
    id         uuid primary key default gen_random_uuid(),
    name       text      not null unique,
    url        text      not null unique,
    user_id    uuid      not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    foreign key (user_id)
        references users (id)
        on delete cascade
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table feeds;
-- +goose StatementEnd
