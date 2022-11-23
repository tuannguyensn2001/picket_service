create table if not exists tests
(
    id                   serial primary key,
    code                 varchar(100),
    name                 varchar(255),
    time_to_do           int,
    time_start           timestamp,
    time_end             timestamp,
    do_once              boolean,
    password             varchar(50),
    prevent_cheat        smallint,
    is_authenticate_user boolean,
    show_mark            smallint,
    show_answer          smallint,
    version              int,
    created_by           int,
    created_at           timestamp,
    updated_at           timestamp,
    deleted_at           timestamp
)