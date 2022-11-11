CREATE table transaction
(
    transaction_id serial PRIMARY KEY,
    user_id        integer,
    currency_id    integer,
    account_id     integer,
    amount         DECIMAL(19, 4),
    description    varchar(255),
    type           boolean     NOT NULL,
    suspended      bool                 DEFAULT false,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES account (account_id),
    FOREIGN KEY (currency_id) REFERENCES currency (currency_id)
)