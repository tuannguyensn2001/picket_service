create table test_multiple_choice_answers (
    id serial primary key ,
    test_multiple_choice_id int,
    answer varchar(50),
    score float,
    type smallint,
    created_at timestamp,
    updated_at timestamp
)