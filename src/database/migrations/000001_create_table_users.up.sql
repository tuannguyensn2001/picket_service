create table users
(
    id         int     not null auto_increment,
    username   varchar(255),
    email      varchar(255) not null,
    password   varchar(255),
    type       int,
    status     int default 1,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    primary key (id)
)