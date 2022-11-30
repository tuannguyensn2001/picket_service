create table test_multiple_choice
(
    id         serial primary key,
    file_path  varchar(255),
    score      int,
    created_at timestamp,
    updated_at timestamp
)