CREATE table currency
(
    currency_id serial PRIMARY KEY,
    name        varchar(50) not null UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO currency (name, created_at, updated_at)
VALUES ('$', NOW(), NOW());

INSERT INTO currency (name, created_at, updated_at)
VALUES ('€', NOW(), NOW());

INSERT INTO currency (name, created_at, updated_at)
VALUES ('₺', NOW(), NOW());

INSERT INTO currency (name, created_at, updated_at)
VALUES ('₽', NOW(), NOW());