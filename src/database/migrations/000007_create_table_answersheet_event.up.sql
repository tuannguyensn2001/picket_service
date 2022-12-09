create table answersheet_event(
    id serial primary key ,
    user_id int,
    test_id int,
    event varchar(50),
    created_at timestamp,
    updated_at timestamp
)