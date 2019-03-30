-- Your SQL goes here
CREATE TABLE account_pending_emails (
    id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    version BIGINT NOT NULL,

    email TEXT NOT NULL,
    token TEXT NOT NULL,
    trials BIGINT NOT NULL,

    account_id UUID NOT NULL REFERENCES account_accounts (id),

    PRIMARY KEY(id)
);

CREATE TABLE account_pending_emails_events (
    id UUID NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    aggregate_id UUID NOT NULL REFERENCES account_pending_emails (id),
    data JSONB NOT NULL,
    metadata JSONB NOT NULL,

    PRIMARY KEY(id)
);
