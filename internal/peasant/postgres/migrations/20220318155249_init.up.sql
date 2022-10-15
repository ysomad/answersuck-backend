CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    email varchar(320) UNIQUE NOT NULL,
    username varchar(24) UNIQUE NOT NULL,
    password text NOT NULL,
    is_email_verified boolean DEFAULT FALSE NOT NULL,
    is_archived boolean DEFAULT FALSE NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    updated_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS email_verification (
    account_id uuid NOT NULL REFERENCES account (id),
    code varchar(255) UNIQUE NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL
);