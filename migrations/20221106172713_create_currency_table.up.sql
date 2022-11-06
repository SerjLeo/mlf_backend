CREATE table currencies (
    currency_id  serial PRIMARY KEY,
    name         varchar(50) not null UNIQUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
)