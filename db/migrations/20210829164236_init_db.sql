-- +goose Up
create table if not exists recipe (
    id bigserial primary key,
    user_id bigint,
    name text not null unique,
    description text not null,
    actions text[]
);

-- +goose Down
drop table recipe;