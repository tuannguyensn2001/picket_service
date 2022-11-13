create table members
(
    id         int not null auto_increment,
    user_id    int,
    class_id   int,
    role       int,
    status     int,
    created_at timestamp,
    updated_at timestamp,
    primary key (id)
)