CREATE table users
(
    user_id      serial PRIMARY KEY,
    name         varchar(50),
    email        varchar(50) not null,
    hashed_pass  varchar(50) not null,
    is_confirmed BOOLEAN DEFAULT false,
    user_role    integer DEFAULT 1
)