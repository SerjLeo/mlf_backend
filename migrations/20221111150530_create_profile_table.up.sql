CREATE table profile
(
    profile_id  serial PRIMARY KEY,
    name        varchar(50),
    currency_id integer,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (currency_id) REFERENCES currency (currency_id)
);