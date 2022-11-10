create table classes
(
    id          int not null auto_increment,
    name        varchar(255),
    description varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp,
    primary key (id)
)