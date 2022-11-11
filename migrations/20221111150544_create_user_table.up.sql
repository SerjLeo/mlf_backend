CREATE table users
(
    user_id      serial PRIMARY KEY,
    email        varchar(50) not null UNIQUE,
    hashed_pass  varchar(60) not null,
    currency_id  integer              DEFAULT 1,
    profile_id   integer,
    is_confirmed BOOLEAN              DEFAULT false,
    user_role    integer              DEFAULT 1,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (currency_id) REFERENCES currency (currency_id),
    FOREIGN KEY (profile_id) REFERENCES profile (profile_id)
);