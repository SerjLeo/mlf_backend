CREATE table account_balance
(
    account_balance_id serial PRIMARY KEY,
    CONSTRAINT account_balance_id PRIMARY KEY (account_id, balance_id),
    account_id         integer,
    balance_id         integer,
    suspended          bool                 DEFAULT false,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES account (account_id),
    FOREIGN KEY (balance_id) REFERENCES balance (balance_id)
);