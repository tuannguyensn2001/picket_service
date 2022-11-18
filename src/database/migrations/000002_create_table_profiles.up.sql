create table if not exists profiles
(
    id         serial primary key,
    user_id    int,
    avatar     varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)