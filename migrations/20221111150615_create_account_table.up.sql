CREATE table account
(
    account_id serial PRIMARY KEY,
    name       VARCHAR(255),
    user_id    integer,
    suspended  bool                 DEFAULT false,
    is_default bool                 DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);