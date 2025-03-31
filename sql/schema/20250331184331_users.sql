-- +goose Up
-- +goose StatementBegin
create table users
(
    id         uuid default gen_random_uuid(),
    name       varchar(255) not null,
    created_at timestamp    not null,
    updated_at timestamp    not null
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
