CREATE table transaction
(
    transaction_id serial PRIMARY KEY,
    user_id        integer,
    amount         DECIMAL(8, 2),
    description    varchar(255),
    type           boolean     NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
)