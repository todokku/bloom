-- Your SQL goes here
CREATE TABLE drive_files (
    id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    version BIGINT NOT NULL,

    name TEXT NOT NULL,
    type TEXT NOT NULL, -- MIME type
    size BIGINT NOT NULL,
    parent_id UUID REFERENCES drive_files (id),
    removed_at TIMESTAMP WITH TIME ZONE,

    owner_id UUID NOT NULL REFERENCES kernel_accounts (id),

    PRIMARY KEY(id)
);

CREATE TABLE drive_files_events (
    id UUID NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    aggregate_id UUID NOT NULL REFERENCES drive_files (id),
    data JSONB NOT NULL,
    metadata JSONB NOT NULL,

    PRIMARY KEY(id)
);
