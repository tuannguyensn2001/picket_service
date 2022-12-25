create table jobs
(
    id         serial primary key,
    payload    json,
    status     varchar(100),
    error_message varchar(255),
    created_at timestamp,
    updated_at timestamp
)