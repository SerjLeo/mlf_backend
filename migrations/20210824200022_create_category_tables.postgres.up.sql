CREATE table category
(
    category_id serial PRIMARY KEY,
    user_id     integer,
    name        varchar(50) NOT NULL UNIQUE,
    color       varchar(7)  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
)