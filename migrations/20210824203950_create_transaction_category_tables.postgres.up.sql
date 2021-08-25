CREATE table transaction_category
(
    transaction_category_id serial PRIMARY KEY,
    user_id                 integer,
    category_id             integer,
    transaction_id          integer,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (category_id) REFERENCES category (category_id),
    FOREIGN KEY (transaction_id) REFERENCES transaction (transaction_id)
)