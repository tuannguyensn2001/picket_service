create table profiles
(
    id         int not null auto_increment,
    user_id    int,
    avatar     varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    primary key (id)
)