
-- +migrate Up
CREATE TABLE IF NOT EXISTS slack_channels (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    channel_id text NOT NULL CONSTRAINT slack_channel_id_uq UNIQUE, 
    team_id uuid NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL   
);

CREATE INDEX slack_channel_team_id_idx ON slack_channels (team_id);

-- +migrate Down
DROP INDEX IF EXISTS slack_channel_team_id_idx;
DROP TABLE IF EXISTS slack_channels;