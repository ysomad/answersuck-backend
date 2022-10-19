CREATE TABLE IF NOT EXISTS password_token (
    account_id uuid NOT NULL REFERENCES account (id),
    token char(128) UNIQUE NOT NULL,
    expires_at timestamptz NOT NULL
);