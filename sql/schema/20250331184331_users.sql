-- +goose Up
-- +goose StatementBegin
create table users
(
    id         uuid primary key default gen_random_uuid(),
    name       text not null unique,
    created_at timestamp    not null,
    updated_at timestamp    not null
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
