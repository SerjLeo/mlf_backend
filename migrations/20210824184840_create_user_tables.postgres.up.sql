CREATE table users
(
    user_id      serial PRIMARY KEY,
    name         varchar(50),
    email        varchar(50) not null UNIQUE,
    hashed_pass  varchar(60) not null,
    is_confirmed BOOLEAN              DEFAULT false,
    user_role    integer              DEFAULT 1,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
)