CREATE table account_balance
(
    CONSTRAINT account_balance_id PRIMARY KEY (account_id, balance_id),
    account_id         integer,
    balance_id         integer,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES account (account_id),
    FOREIGN KEY (balance_id) REFERENCES balance (balance_id)
);