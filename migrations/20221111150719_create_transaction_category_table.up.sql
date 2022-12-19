CREATE table transaction_category
(
    user_id        integer,
    category_id    integer,
    transaction_id integer,
    CONSTRAINT transaction_category_id PRIMARY KEY (category_id, transaction_id),
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category (category_id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_id) REFERENCES transaction (transaction_id) ON DELETE CASCADE
)