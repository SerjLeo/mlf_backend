CREATE table balance
(
    balance_id  serial PRIMARY KEY,
    amount      DECIMAL(19, 4) DEFAULT 0,
    currency_id integer,
    user_id     integer,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (currency_id) REFERENCES currency (currency_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);